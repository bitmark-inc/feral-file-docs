# Feral File — Artist Installation Guide
## Artist Installation FAQ

- Questions from Artist: Can we embed a self-hosted website on Feral File? For example, how is [Integrity Protocol](https://feralfile.com/explore/series/integrity-protocol-lower-limb-lecture-cop) embedded in the current show? Are you saying embedding is done by uploading the entire website in a folder?
- Answer to reply: We do not support embedding a self-hosted website. Instead, please refer to the attached images of what [American Artist and Tommy Martinez](https://feralfile.com/explore/artists/american-artist-and-tommy-martinez-plp) uploaded in the source file. A file named "index.html" is required in the folder for proper installation.
---

- Questions from Artist: When you say, "Your edition can’t require any external resources or servers to operate," is this a policy decision or a technical limitation? Is the Feral File framework literally unable to access an external script or server at runtime?
- Answer to reply: It’s not really a technical limitation. We do that for a number of reasons. 
1. We allow the user to download the full source code bundle. If it loads the external script, when the host of that script dies, users will no longer be able to open the artwork on their local.
2. Platforms like Feral File, OpenSea, Objkt will need to show this artwork on the website. We generally prefer the artwork to be self-contained so that:
   2.1. The artwork can be stored on IPFS, making the artwork permanently available for viewing. This is very important for NFT collectors.
   2.2. External links can introduce security risks and trust issues, as the content on the external source can change (intentionally change, or get hacked), contradicting the immutable nature of NFTs, making (part of) its content not decentralized.
---

- Questions from Artist: In your example, the filenames are listed as `filename_1`, `filename_2`, etc. Do these filenames have any restrictions? Can I use underscores in my filename? For example, can my file be named `A_B_1`? For the artwork folder, your example says to name the main video file `My_video.mov`. Do I have to name it `My_Video`, or can it have a different filename?
- Answer to reply: The system likely splits off the final underscore character from the filename string to extract the number. Therefore, other underscores earlier in the filename should work fine. `My_Video` is just an example. You should be able to name the file whatever you like. 
---

- Questions from Artist: I'm trying to upload a ZIP of an `.app` executable. I'm still finishing the software (will include a Windows version and README as well), but I noticed whenever I tried to upload I would get errors on Feral File telling me it was an invalid format. I would also like to specify images for each individual edition. How do I do that if the work is "software" or should I categorize it as image, and zip the app executable with each individual image? Thank you!
- Answer to reply: I think the best solution for our system is to choose "image" as your type and then describe the actual medium using the "+ Add custom medium" option. This should allow you to upload your custom images for each artwork in the series while also clearly communicating to collectors what the actual artwork is. As you've probably noticed, if you select the "software" option, the system expects the file to be p5.js, GLSL, HTML, or CSS.
---

- Questions from Artist:
- Answer to reply:
---
