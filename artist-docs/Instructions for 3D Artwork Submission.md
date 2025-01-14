# Feral File — Artist Installation Guide
## Artist Installation Guide Supported File Formats
### Instructions for 3D Artwork Submission

Examples: [Sophie Khan](https://feralfile.com/artworks/spectrality-from-the-divers-ix-kkw), [Auriea Harvey](https://feralfile.com/exhibitions/peer-to-peer-pjb/series/auriea-harvey-9kg?viewMode=Grid), [Carla Gannis](https://feralfile.com/artworks/peep-o-rama-r4s)

The <modelviewer> tag and a GLB file are the default way to create a 3D edition for Feral File. If you use the URL [https://modelviewer.dev/editor/](https://modelviewer.dev/editor/) to create your files, things will look like this: 

📁 my_artwork.zip
* index.html (generated from the link above, required at the root of the ZIP)
* [model-viewer.min.js](https://unpkg.com/@google/model-viewer@1.7.2/dist/model-viewer.min.js) (very important! all files must be inside the ZIP and not linking out to other sites online. this isn’t downloaded by default. click the link to get the file)
* script.js (generated from the link above)
* styles.css (generated from the link above)
* snippet.txt (generated from the link above)
* my_model.glb
* maybe_a_PDF_with_other_information.pdf
* maybe_a_README_file_with_other_information.txt
* 📁 maybe_other? (if you want to include your model in other formats)
    * my_model.stl
    * my_model.usdz
    * my_model.dae
    * my_model.obj
    * etc.
* 📁 maybe_a_video?
    * my_artwork.mov

Your edition can’t require any external resources or servers to operate. The entire work must be contained within this ZIP. 

***** When you create the ZIP, select the files and make the ZIP; don’t ZIP the folder that encloses the files. After the ZIP is created, you can rename it.

***** Sometimes the Feral File code can’t recognize this type of file automatically, so the MIME type needs to be set manually by Bitmark. If your artwork isn’t working a few minutes after uploading, please tag @rita in your exhibition channel on Discord. This only takes a few minutes, but because of the time zone, it sometimes takes hours for the engineers to see the request.

If you need help with 3D model format, you can email Oren directly ([oren.avraham.shoham@gmail.com](mailto:oren.avraham.shoham@gmail.com)).

This is Oren's fee structure:

For a 3D viewer website made with [https://modelviewer.dev](https://modelviewer.dev/) displaying a single model (similar to what I set up for Peter's COUNT ZERO FOUR piece), I'd ask for a $500 flat fee, which would include helping the artist convert their model to GLB format and configuring the &lt;model-viewer> camera and controls. If an artist wanted something more complex than what &lt;model-viewer> affords, e.g. a site with custom HTML/CSS, custom interaction logic, multiple 3D models, or JavaScript-based animation, or if they needed help with a pre-existing 3D viewer setup, I'd charge an hourly rate of $60.
