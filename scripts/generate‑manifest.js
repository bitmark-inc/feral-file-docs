const { readdir, readFile, writeFile } = require('fs').promises;
const path  = require('path');


const LEARN_DIR = path.join(__dirname, '..', 'learning');
const OUT_FILE = path.join(LEARN_DIR, 'learn-manifest.json');

function parseFrontMatter(txt) {
  const m = txt.match(/^---\s*([\s\S]*?)\s*---/);
  if (!m) return {};
  const obj = {};
  m[1].split('\n').forEach(line => {
    const [k, ...rest] = line.split(':');
    if (k && rest.length) {
      obj[k.trim()] = rest.join(':').trim().replace(/^"(.*)"$/, '$1');
    }
  });
  return obj;
}

(async () => {
  // read all markdown files except guidelines.md
  const files = (await readdir(LEARN_DIR))
    .filter(f => f.endsWith('.md') && f !== 'guidelines.md');

  const articles = [];

  for (const file of files) {
    const raw = await readFile(path.join(LEARN_DIR, file), 'utf8');
    const fm  = parseFrontMatter(raw);
    if (fm.draft === 'true') continue;              // skip drafts

    articles.push({
      slug: file.replace(/\.md$/, ''),
      navTitle: fm.navTitle || fm.title,
      title:  fm.title,
      description: fm.description,
      weight: Number(fm.weight) || 9999
    });
  }

  articles.sort((a, b) => a.weight - b.weight);

  await writeFile(
    OUT_FILE,
    JSON.stringify(
      { generated: new Date().toISOString(), articles },
      null,
      2
    )
  );

  console.log(`✅  Manifest written → ${OUT_FILE} (${articles.length} articles)`);
})();
