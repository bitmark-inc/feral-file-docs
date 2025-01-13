# Feral File — Artist Installation Guide

## Artist Flow: Artwork Installation Process

This guide outlines the steps artists need to take to successfully display their artwork on Feral File. Follow the instructions below to ensure a smooth and seamless installation process.

1. Provide Required Information

To begin, the artist must provide the following details to the curator or Rita (from the Feral File team):

- **Email Address**
- **Legal Name**
- **Wallet Address**

This information will be entered into the admin system to grant the artist access to the Feral File platform.

2. Receive Access to the Admin System

The artist will receive an email from **[support@feralfile.com](mailto:support@feralfile.com)** containing a link to access the Feral File admin portal.

- Click the link and log in using the wallet address provided to the curator or Rita.

3. Update Your Profile

After logging into the admin portal, navigate to the **Profile Page** and update your information:

3.1. **Upload Your Artist Avatar**
   - **File guidelines:**
     - Formats: PNG, JPG, JPEG
     - Maximum file size: 2MB
   - You can view other artist profiles on the platform for inspiration. [See what others have uploaded](https://feralfile.com/explore/artists?&sortBy=recent).

3.2. **Review and Update Additional Profile Details**
   - Add your **location**, **website**, **Instagram handle**, and **Twitter handle**.
   - Submit your biography to your curator; it will be copy-edited by Feral File's editor and uploaded on your behalf.

3.4. Sign the Artist Agreement

Once your profile is updated:

- You will be prompted to review and sign the **Artist Agreement**. This step is mandatory to proceed further and begin the installation process.

3.5. Upload Your Artwork

Once your profile is complete, you can begin uploading your artwork.

Required Files:

1. **Thumbnail Image**
   - This image is displayed on the [index for the series](https://feralfile.com/exhibitions/graph-8vs/series/fictional-lullaby-pe1?viewMode=Grid) and when people [browse the work](https://feralfile.com/exhibitions/graph-8vs?tab=Collect#1) on the “overview” page.
   - **File specifications:**
     - Supported formats: PNG, JPG, JPEG
     - Maximum file size: 11MB

2. **Main Artwork File**
   - The format and specifications for this file will depend on the type of artwork being uploaded (e.g., image, video, generative art).
   - Refer to the specific **Supported File Formats** section for detailed requirements.

---

Additional Notes

- Ensure all files meet the size and format specifications to avoid upload errors.
- If you encounter any issues during the installation process, contact **[support@feralfile.com](mailto:support@feralfile.com)** for assistance.

By following these steps, artists can ensure their work is accurately and professionally displayed on the Feral File platform.

## Supported File Formats

### 1. Instructions for Web/Software Artwork Submission

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

### 2. Instructions for Video Artwork Submission

* **Recommendation: ProRes 422 or 422 LT, MP4 (use H.264 video encoding to maximize the browser compatibility).**
* Possible: ProRes 422 or 422 LT, MP4 (H.264), WebM, MKV, AVI, MOV, WMV
* You can choose to “Play video in loop” a video, it won’t loop by default.
* Your video can be 4K or HD and you can include both in the download package.

### Video Upload Details

Examples: [Holly Herndon and Mat Dryhurst](https://feralfile.com/exhibitions/sound-machines-xz1/series/play-from-memory-xyq?viewMode=Grid), [Lia Coleman](https://feralfile.com/exhibitions/n-12-2ts/series/emergence-nji?viewMode=Grid)

We require the artwork files for the following sections:

1. Thumbnails: PNG, JPG, JPEG (max. file size 11MB per file) The thumbnail will appear on the [Artwork detail page](https://feralfile.com/exhibitions/sound-machines-xz1/series/play-from-memory-xyq?viewMode=Grid). *Format example: [Thumbnail folder](https://drive.google.com/drive/folders/14DJVCwD0bfc5TRhZi6CL9XLUCZiFz4Y-?usp=share_link)*

📁 artwork_thumbnail
* filename_1.jpg (#1)
* filename_2.jpg (#2)
* filename_3.jpg (#3)
* filename_4.jpg (#4)

We recommend creating two versions of your video for Feral File. One version is for web streaming and the other version is the highest resolution you have to be downloaded by the collector.

*If you want to have the MP4 for streaming and a high-resolution version, and other materials for the collector, you’ll upload the MP4 and separately upload a ZIP file for the collector to download. Here’s an example ZIP file format.*

2. Web streaming: The compressed version of your video for web steaming must be no greater than 500MB per file, and it must be a MP4. (This is the video that users will see when they click “[Artwork](https://feralfile.com/exhibitions/sound-machines-xz1/series/play-from-memory-xyq?viewMode=Grid&artPlayTokenIndex=0)”). *Format example: [Web streaming Preview folder](https://drive.google.com/drive/folders/1c5fiEiekG4OBfQWaLlXUVR5WYkztTMKk?usp=share_link)*

📁 artwork_web
* filename_1.mp4 (#1)
* filename_2.mp4 (#2)
* filename_3.mp4 (#3)
* filename_4.mp4 (#4)

***** The compressed version of your video for web steaming must be no greater than 500 MB, and it must be an MP4.

3. Source file: This is the highest-quality file of your artwork. You can upload a <span style="text-decoration:underline;">single video file</span> (in MP4 or MOV or WMV or QUICKTIME or AVI or WEBM or MKV) or a <span style="text-decoration:underline;">ZIP package</span> that includes the video file and other documentation, such as text, process images, etc.  you want the collector to have. (max. file size 5GB per file) _Format example: [Artwork folder](https://drive.google.com/drive/folders/1yav-o_uUiy6bGtKbY87d7NdddOdSlAm9?usp=share_link)_

📁 artwork
 
   📁 filename_1.zip (Folder 1 is for Edition 1) or filename_1.mp4 (#1)
   * My_video.mov (required at the root of the ZIP, highest quality possible for the collector)
   * my_video_alt.mov (you can include another video too if you want)
   * 1.jpg
   * maybe_a_PDF_with_other_information.pdf
   * maybe_a_README_file_with_other_information.txt
   * (anything else you want the collector to have)
  
   📁 filename_2.zip (Folder 2 is for Edition 2) or filename_2.mp4 (#2)
   * my_video.mov (required at the root of the ZIP, highest quality possible for the collector)
   * my_video_alt.mov (you can include another video too if you want)
   * 2.jpg
   * maybe_a_PDF_with_other_information.pdf
   * maybe_a_README_file_with_other_information.txt
   * (anything else you want the collector to have)

### 3. Instructions for 3D Artwork Submission

* We’ve decided to make 3D uploads the same as the Software category. This offers each artist the most flexibility in how their work is displayed. Upload a ZIP file with an index.html at the top of the ZIP. See the “3D Upload Details” section below. 
* Include all 3D files you want the collector to have access to within the ZIP
* Include other information, such as text, process images, etc. you want the collector to have within the ZIP. The ZIP upload has a 5GB limit.
* The <modelviewer> tag is the a default way to create a 3D edition
    * The basics: https://modelviewer.dev/
    * Explore examples: https://modelviewer.dev/examples/loading/index.html
    * Generate your files: https://modelviewer.dev/editor/
* Feral File can wrap these files into Electron applications to allow the collector to download Linux, Windows, and Mac OS applications to view your files on their computers.

### 3D Upload Details

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

### 4. Instructions for Text Artwork Submission

Examples: [Martin John Callanan](https://feralfile.com/artworks/each-and-every-command-gmb?sort=0)

We require the artwork files for the following sections:

1. Thumbnails: PNG, JPG, JPEG (max. file size 11MB per file) _Format example: [Thumbnail folder](https://drive.google.com/drive/folders/14DJVCwD0bfc5TRhZi6CL9XLUCZiFz4Y-?usp=share_link)_

📁 artwork_thumbnail
* filename_1.jpg (#1)
* filename_2.jpg (#2)
* filename_3.jpg (#3)
* filename_4.jpg (#4)

2. Web streaming: Upload your txt as a single txt file that will work for web viewing (max. file size 5GB per file) _Format example: [Web streaming Preview folder](https://drive.google.com/drive/folders/1c5fiEiekG4OBfQWaLlXUVR5WYkztTMKk?usp=share_link)_

📁 artwork_web
* filename_1.txt (#1)
* filename_2.txt  (#2)
* filename_3.txt  (#3)
* filename_4.txt  (#4)

3. Source file: You can upload a <span style="text-decoration:underline;">single TXT file</span> or a <span style="text-decoration:underline;">ZIP packag</span>e that includes the TXT file and other documentation, such as text, process images, etc.  you want the collector to have. (max. file size 5GB per file) _Format example: [Artwork folder](https://drive.google.com/drive/folders/1yav-o_uUiy6bGtKbY87d7NdddOdSlAm9?usp=share_link)_

📁 my_artwork_name.zip
* my_artwork.txt (required at the root of the ZIP)
* maybe_a_PDF_with_other_information.pdf

### 5. Instructions for Animated GIF Artwork Submission

We require the artwork files for the following sections:

1. Thumbnails: PNG, JPG, JPEG (max. file size 11MB per file) _Format example: [Thumbnail folder](https://drive.google.com/drive/folders/14DJVCwD0bfc5TRhZi6CL9XLUCZiFz4Y-?usp=share_link)_

📁 artwork_thumbnail
* filename_1.jpg (#1)
* filename_2.jpg (#2)
* filename_3.jpg (#3)
* filename_4.jpg (#4)

2. Web streaming: Upload your animated GIF as a single GIF file that will work for web viewing (max. file size 22MB per file) _Format example: [Web streaming Preview folder](https://drive.google.com/drive/folders/1c5fiEiekG4OBfQWaLlXUVR5WYkztTMKk?usp=share_link)_

📁 artwork_web
* filename_1.gif (#1)
* filename_2.gif  (#2)
* filename_3.gif  (#3)
* filename_4.gif  (#4)

3. Source file: You can upload a <span style="text-decoration:underline;">single animated GIF file</span> or a <span style="text-decoration:underline;">ZIP package</span> that includes the animated GIF file and other documentation, such as text, process images, etc. you want the collector to have. (max. file size 5GB per file) _Format example: [Artwork folder](https://drive.google.com/drive/folders/1yav-o_uUiy6bGtKbY87d7NdddOdSlAm9?usp=share_link)_

📁 my_artwork_name.zip
* my_artwork.gif (required at the root of the ZIP)
* maybe_a_PDF_with_other_information.pdf

### 6. Instructions for Audio Artwork Submission

* __Recommendation: WAV__
* Possible: WAV, AAC, AIFF, MP3, WMA, M4A, AMR, FLAC, MID, OGG

We require the artwork files for the following sections:

1. Thumbnails: PNG, JPG, JPEG (max. file size 11MB per file) _Format example: [Thumbnail folder](https://drive.google.com/drive/folders/14DJVCwD0bfc5TRhZi6CL9XLUCZiFz4Y-?usp=share_link)_

📁 artwork_thumbnail
* filename_1.jpg (#1)
* filename_2.jpg (#2)
* filename_3.jpg (#3)
* filename_4.jpg (#4)

2. Web streaming: Upload your audio as a single M4A file that will work for web viewing (max. file size 110MB per file) _Format example: [Web streaming Preview folder](https://drive.google.com/drive/folders/1c5fiEiekG4OBfQWaLlXUVR5WYkztTMKk?usp=share_link)_

 📁 artwork_web
* filename_1.m4a (#1)
* filename_2.m4a (#2)
* filename_3.m4a (#3)
* filename_4.m4a (#4)

3. Source file: You can upload a <span style="text-decoration:underline;">single audio file</span> or a <span style="text-decoration:underline;">ZIP package</span> that includes the audio file and other documentation, such as text, process images, etc. you want the collector to have. (max. file size 5GB per file) _Format example: [Artwork folder](https://drive.google.com/drive/folders/1yav-o_uUiy6bGtKbY87d7NdddOdSlAm9?usp=share_link)_

📁 my_artwork_name.zip
* my_artwork.wav (required at the root of the ZIP)
* maybe_a_PDF_with_other_information.pdf

### 7. Instructions for Image Artwork Submission
* __Recommendation: PNG__
* Possible: PNG, JPEG, JPG, BMP

Examples: [Linda Dounia](https://feralfile.com/exhibitions/in-visible-419/series/chez-jo-ei5?viewMode=Grid)

Feral File requires the artwork files for the following sections:

1. Thumbnails: PNG, JPG, JPEG (max. file size 11MB) _Format example: [Thumbnail folder](https://drive.google.com/drive/folders/14DJVCwD0bfc5TRhZi6CL9XLUCZiFz4Y-?usp=share_link)_

📁 artwork_thumbnail
* filename_1.jpg (#1)
* filename_2.jpg (#2)
* filename_3.jpg (#3)
* filename_4.jpg (#4)

2. Web streaming: The image for web steaming must be no greater than 22MB and it must be a PNG, JPG or JPEG. _Format example: [Web streaming Preview folder](https://drive.google.com/drive/folders/1c5fiEiekG4OBfQWaLlXUVR5WYkztTMKk?usp=share_link)_

📁 artwork_web
* filename_1.jpg (#1)
* filename_2.jpg (#2)
* filename_3.jpg (#3)
* filename_4.jpg (#4)

3. Source file: This is the highest-quality file of your artwork. You can upload a <span style="text-decoration:underline;">single image file</span> or a <span style="text-decoration:underline;">ZIP packag</span>e that includes the image file and other documentation, such as text, process images, etc.  you want the collector to have. (max. file size 5GB per file) _Format example: [Artwork folder](https://drive.google.com/drive/folders/1yav-o_uUiy6bGtKbY87d7NdddOdSlAm9?usp=share_link)_

📁 artwork

📁 filename_1.zip (Folder 1 is for Edition 1) or filename_1.jpg
* 1.jpg
* maybe_a_PDF_with_other_information.pdf

📁 filename_2.zip (Folder 2 is for Edition 2) or filename_2.jpg
* 2.jpg
* maybe_a_PDF_with_other_information.pdf

### 8. Instructions for PDF Artwork Submission

* You can upload an edition as a PDF file
    * *Are there PDF variations, version numbers, etc. that we should be aware of that might create issues?*

If your artwork format is PDF, please reach out to [rita@feralfile.com](mailto:rita@feralfile.com) before the installation, so we can make sure the system can support your files.
