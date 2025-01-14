# Feral File — Artist Installation Guide
## Artist Installation Guide Supported File Formats
### Instructions for Web/Software Artwork Submission

* Upload as a ZIP. Anything that runs in a browser should run, typically HTML/CSS/JavaScript.
* A file named “index.html” is required. It will load first and it needs to be in the *root folder/directory* of the ZIP. Please see the Web/Software Upload Details section below for the details.
* Include all additional files you want the collector to have, such as text and other media within the ZIP. The ZIP upload has a 5GB limit.
* All files should be self-contained within the ZIP file. It’s not possible to set up a server or database to support media on Feral File, and a network connection shouldn’t be required to run the work after it’s downloaded. Do not link to external assets. 
* Considerations
    * Resolution — How does your work scale? You should feel good about how the software looks at different pixel resolutions, for example all the way from a mobile phone to 4K (3840 × 2160). __Test this!__
    * Cursor — What about the cursor? Do you want it to be there or to hide it? Or, to put it on a timer to disappear after a while and reappear when the mouse is moved?  
    * Ratio — How does your work respond to different aspect ratios and vertical vs. horizontal? 
    * Frame Rate — What is the ideal frame rate for your software? 24? 30? 60? Computers get faster over time, so set a limit if you want your software to remain at a specific frame rate over time.
    * Resize — The work should adjust size when the browser is resized or the application is run on different screens. For example, when your work goes into full screen mode, it should adjust to use all the pixels. When the work leaves full screen mode, it should adjust for the current browser size. We can work on templates together to share code to do this.
    * License — Consider adding a software license to your source code to clarify the status of the code to the Collector. [This Wikipedia page outlines the basics](https://en.wikipedia.org/wiki/Software_license).

### Web/Software Upload Details

Artwork library documentation: [https://github.com/bitmark-inc/feral-file-docs/blob/master/artist-docs/feralfile-artwork-library.md](https://github.com/bitmark-inc/feral-file-docs/blob/master/artist-docs/feralfile-artwork-library.md) 

Examples: [Elsif](https://feralfile.com/exhibitions/n-12-2ts/series/superbloom-1mm?viewMode=Grid), [Piter Pasma](https://feralfile.com/exhibitions/n-12-2ts/series/balls-to-the-walls-zdn?viewMode=Grid), [Anna Lucia](https://feralfile.com/exhibitions/n-12-2ts/series/oscillation-ii-0wy?viewMode=Grid)

Feral File supports HTML/CSS/Javascript as a “media type” for exhibitions. To upload this type of work, create a single ZIP file to contain all of the files needed to run. Your edition can’t require any external resources or servers to operate. The entire work must be contained within this ZIP. 

The main HTML file must be named “index.html” and must be in the root folder of the ZIP. The name of the folder and names of the other files can be anything you like. For example:

📁 my_edition.zip
* index.html
* style.css
* sketch.js
* p5.js
* my_artwork.png

You can also have subfolders, so you can format like this too if you prefer, but “index.html” needs to be at the top:

📁 my_edition.zip
* index.html
* 📁 css
    * style.css
* sketch.js
* p5.js
* my_artwork.png

***** When you create the ZIP, select the files and make the ZIP; don’t ZIP the folder that encloses the files. After the ZIP is created, you can rename it.

If you’re working with your own custom code, we’re figuring out other software formats to support on Feral File and we welcome suggestions.

**CSS TEMPLATE**
```
html, body {

  margin: 0;

  padding: 0;

}

canvas {

  display: block;

}
```

**P5.JS TEMPLATE**
```
function setup() {

   createCanvas(windowWidth, windowHeight);

}

function draw() {

   background(frameCount%255);

   fill(255);

   noStroke();

   circle(width/2, height/2, height);

}

function windowResized() {

   resizeCanvas(windowWidth, windowHeight);

}
```

**RANDOM SEED TEMPLATE**

Please see this example:

[https://editor.p5js.org/REAS/sketches/Hftgmr1s1](https://editor.p5js.org/REAS/sketches/Hftgmr1s1)

The specific Feral File code is here:

// DATA from FERAL FILE
```
const urlSearchParams = new URLSearchParams(window.location.search);

const params = Object.fromEntries(urlSearchParams.entries());

let editionNumber = params.edition_number;

let hash = params.token_id_hash;

And

  r = new RND();

Then replace all of your random() functions with:

let x = r.rb(0.0, 1.0); // for example

And

class RND {

    constructor() {

    this.useA = false;

    let sfc32 = function (uint128Hex) {

      let a = parseInt(uint128Hex.substr(0, 8), 16);

      let b = parseInt(uint128Hex.substr(8, 8), 16);

      let c = parseInt(uint128Hex.substr(16, 8), 16);

      let d = parseInt(uint128Hex.substr(24, 8), 16);

      return function () {

        a |= 0;

        b |= 0;

        c |= 0;

        d |= 0;

        let t = (((a + b) | 0) + d) | 0;

        d = (d + 1) | 0;

        a = b ^ (b >>> 9);

        b = (c + (c &lt;< 3)) | 0;

        c = (c &lt;< 21) | (c >>> 11);

        c = (c + t) | 0;

        return (t >>> 0) / 4294967296;

      };

    };

    this.prngA = new sfc32(hash.substr(2, 32));

    this.prngB = new sfc32(hash.substr(34, 32));

    for (let i = 0; i &lt; 1e6; i += 2) {

      this.prngA();

      this.prngB();

    }

  }

  random_dec() {

    this.useA = !this.useA;

    return this.useA ? this.prngA() : this.prngB();

  }

  rb(a, b) {

    return a + (b - a) * this.random_dec();

  }

}
```

The Feral File team will manually upload thumbnails for the artists. Please provide the thumbnails in the following format.

Thumbnails: PNG, JPG, JPEG (max. file size 11MB per file) The thumbnail will appear on the [Artwork detail page](https://feralfile.com/exhibitions/sound-machines-xz1/series/play-from-memory-xyq?viewMode=Grid). _Format example: [Thumbnail folder](https://drive.google.com/drive/folders/14DJVCwD0bfc5TRhZi6CL9XLUCZiFz4Y-?usp=share_link)_

📁 artwork_thumbnail
* filename_1.jpg (#1)
* filename_2.jpg (#2)
* filename_3.jpg (#3)
* filename_4.jpg (#4)

