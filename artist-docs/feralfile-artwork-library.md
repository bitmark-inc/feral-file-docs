# Feral File Artwork Library

###### tags: `Feral File` `operation` 

## Feral File Variables

For Feral File generative artworks, pre-defined variables are **injected** when the collector clicks to view the software artwork in a frame. This helps artists create variants of their artworks. The variables are:

- `blockchain` - (`ethereum` | `tezos`)
- `contract` - Contract address if any
- `token_id` - A unique ID of a token in the blockchain. It is decimal in Tezos and Ethereum
- `edition_number` - The edition number of the artwork
- `artwork_number` - The artwork number within the series

These variables are passed as URL parameters to the artwork. Here is an example of the parameters passed in the URL to the artwork:

```
https://cdn.feralfileassets.com/previews/ccd387c2-4762-4037-8752-bbc223957199/1727153365/index.html?edition_number=1&artwork_number=2&blockchain=ethereum&contract=0x5D303A17883C442603B8afb54921d1b7DCc0C831&token_id=27351590983062365677486503722577408930064250723080863920865112148460960306241&display_mode=crop
```

### Token ID

By design, token contracts require a unique ID for each token. However, in our exhibition structure, we have multiple series, and each series has its own artwork with an index starting from 0. This can cause duplication of token IDs in a contract if we simply use the index as the token ID. To avoid conflicts, we generate deterministic unique token IDs by combining the series ID and token indexes. The resulting `token_id` is passed as a URL parameter to the artwork and can be used as a random seed.

## Deterministic Randomness

To provide randomness to generative artworks, we offer a random function based on `sfc32`. The function takes `token_id` (passed as a URL parameter) as the seed of the randomness. This ensures that the randomness is deterministic. If the `token_id` is not given, the snippet will create a random one so you can test locally.

### Code

```javascript
<script id="ffrand-snippet">
  function randomTokenId() {
    return Math.floor(Math.random() * Number.MAX_SAFE_INTEGER).toString();
  }

  const urlSearchParams = new URLSearchParams(window.location.search);
  const params = Object.fromEntries(urlSearchParams.entries());
  let token_id = params.token_id || randomTokenId();
  console.log('Token ID:', token_id);

  // Seed generation function
  function cyrb128(str) {
    let h1 = 1779033703, h2 = 3144134277,
        h3 = 1013904242, h4 = 2773480762;
    for (let i = 0, k; i < str.length; i++) {
      k = str.charCodeAt(i);
      h1 = (h2 ^ Math.imul(h1 ^ k, 597399067)) | 0;
      h2 = (h3 ^ Math.imul(h2 ^ k, 2869860233)) | 0;
      h3 = (h4 ^ Math.imul(h3 ^ k, 951274213)) | 0;
      h4 = (h1 ^ Math.imul(h4 ^ k, 2716044179)) | 0;
    }
    h1 = Math.imul(h3 ^ (h1 >>> 18), 597399067) | 0;
    h2 = Math.imul(h4 ^ (h2 >>> 22), 2869860233) | 0;
    h3 = Math.imul(h1 ^ (h3 >>> 17), 951274213) | 0;
    h4 = Math.imul(h2 ^ (h4 >>> 19), 2716044179) | 0;
    return [(h1^h2^h3^h4) >>> 0, h2>>>0, h3>>>0, h4>>>0];
  }

  // sfc32 random number generator
  function sfc32(a, b, c, d) {
    return function() {
      a >>>= 0; b >>>= 0; c >>>= 0; d >>>= 0; 
      var t = (a + b | 0) + d | 0;
      d = d + 1 | 0;
      a = b ^ b >>> 9;
      b = c + (c << 3) | 0;
      c = (c << 21 | c >>> 11);
      c = c + t | 0;
      return (t >>> 0) / 4294967296;
    }
  }

  // Create seed and random function
  const seed = cyrb128(token_id);
  var ffrand = sfc32(seed[0], seed[1], seed[2], seed[3]);
</script>
```

### Usage

After loading the snippet, call `ffrand()` to get a random number.

```javascript
<script>
  let r = ffrand();
  console.log('Random number:', r);
</script>
```

:::warning
When designing your randomness, ensure it remains deterministic based on the given `token_id`; otherwise, the artwork may change with each page reload.
:::

## Artwork Library

With these parameters, the Feral File library retrieves provenance from our indexer, enabling software artwork developers to create variants by integrating this library. Inspired by fxhash, this library is designed as a simple snippet that can be included at the top of generative artworks.

### Functionality

Supported functions:

- Get blockchain height (integer)
- Get provenance (array)

### Usage

Include the script in the `<head>` section and call `ffinit()` when the page loads.

```html
<!DOCTYPE html>
<html lang="en">
<head>
  <script id="ff-library" 
  src="https://ipfs.bitmark.com/ipfs/QmRfJutL1FyAKqT2vAAPErgcJqTfCyWvQtnbiShZhTkNL4"
  type="text/javascript"></script>
</head>

<body onload="ffinit()">
  <!-- Your artwork code here -->
</body>

</html>
```

You can listen to these events in your script:

- `provenance-request-error`
- `provenance-ready`
- `blockchain-info-request-error`
- `blockchain-info-ready`

Example:

```javascript
<script>
  // Listen to blockchain information data
  window.addEventListener("blockchain-info-ready", function (event) {
    console.log("Current block height:", event.detail.height);
  });

  window.addEventListener("blockchain-info-request-error", function (event) {
    console.error(event.detail.error);
  });

  // Listen to token provenance data
  window.addEventListener("provenance-ready", function (event) {
    console.log("Provenance length:", event.detail.provenances.length);
  });

  window.addEventListener("provenance-request-error", function (event) {
    console.error("Failed to get provenance:", event.detail.error);
  });
</script>
```

### Provenance Data Example

```json
[
    {
        "type": "transfer",
        "owner": "KT1GbyoDi7H1sfXmimXpptZJuCdHMh66WS9u",
        "blockchain": "tezos",
        "blockNumber": 2498727,
        "timestamp": "2022-07-02T05:29:14Z",
        "txid": "ooHk3jYeGuSM2ruCPUog5QprLqphjDrNnDNUBoY2UZDWbm14Unf",
        "txURL": "https://tzkt.io/ooHk3jYeGuSM2ruCPUog5QprLqphjDrNnDNUBoY2UZDWbm14Unf"
    },
    {
        "type": "mint",
        "owner": "tz1Ne3XbbFpWAZjPf8QJy5fbuRWpcwUyg42X",
        "blockNumber": 2498678,
        "timestamp": "2022-07-02T05:04:44Z",
        "txid": "oo3X4Fy9RADVk6JVcKS1BeJLjB6bRzwTxLKvbBjFDBdPabHrcmJ",
        "txURL": "https://tzkt.io/oo3X4Fy9RADVk6JVcKS1BeJLjB6bRzwTxLKvbBjFDBdPabHrcmJ"
    }
]
```

## Artwork Attributes

To track artwork attributes, a snippet provides the `$feralfile` namespace with two functions: `$feralfile.features(array_of_attributes)` to input attributes and `$feralfile.getFeatures()` to retrieve saved attributes based on the artwork information.

### Code

```javascript
<script id="feralfile-attributes">
  var $feralfile = {
    attributes: [],
    features: function (attrs) {
      this.attributes = attrs || this.attributes;
    },
    getFeatures: function () {
      return this.attributes || [];
    },
  };
</script>
```

### Usage

Load the snippet and call `$feralfile.features()` to set attributes.

```html
<!DOCTYPE html>
<html lang="en">
<head>
  <script id="ff-attrs" 
  src="https://ipfs.bitmark.com/ipfs/QmakVUXKyHF5cddi6owHLbCib7wbU36nY2JoQiNkE7eSho"
  type="text/javascript"></script>
</head>

<body>
    <script>
      $feralfile.features([
        { name: 'color', value: 'dark' },
        { name: 'tmax', value: 254 },
        { name: 'tmin', value: 189 },
        { name: 'letter', value: 5 },
        { name: 'dotnum', value: 374 },
      ]);
    </script>    
</body>

</html>
```

## Display Mode Adjustment

To ensure artworks display correctly across different devices and screen sizes, we have introduced a `display_mode` URL query parameter. The Feral File apps (mobile and TV) and the Feral File website will include this parameter when loading the artwork. The parameter can have two values: `crop` or `fit`.

Example URL:

```
https://artwork_cdn_url?display_mode=crop
```

### Implementation

Your artwork should read the `display_mode` parameter from the URL and adjust the visual content accordingly:

- **`display_mode=crop`**: The artwork should scale and crop its content to fill the entire viewport, possibly cutting off parts of the content if necessary to maintain aspect ratio.
  
- **`display_mode=fit`**: The artwork should scale its content to fit entirely within the viewport, ensuring that all content is visible, potentially adding padding or black bars if the aspect ratios do not match.

### Code Example

```javascript
<script>
  // Get display_mode from URL parameters
  const urlParams = new URLSearchParams(window.location.search);
  const displayMode = urlParams.get('display_mode') || 'fit'; // Default to 'fit' if not specified

  // Adjust your artwork based on the display mode
  function adjustDisplayMode() {
    const canvas = document.getElementById('myCanvas');
    const context = canvas.getContext('2d');

    const windowWidth = window.innerWidth;
    const windowHeight = window.innerHeight;

    // Set canvas size to match window size
    canvas.width = windowWidth;
    canvas.height = windowHeight;

    if (displayMode === 'crop') {
      // Implement cropping logic to fill the viewport
      // Scale and center your artwork to fill the canvas, possibly cropping edges
      context.save();
      // Your drawing code here, scaling up to fill and possibly cropping
      context.restore();
    } else if (displayMode === 'fit') {
      // Implement fitting logic to ensure all content is visible within the viewport
      // Scale your artwork to fit within the canvas without cropping
      context.save();
      // Your drawing code here, scaling down to fit entirely
      context.restore();
    } else {
      // Handle unexpected values if necessary
      console.warn('Unknown display_mode:', displayMode);
    }
  }

  // Call adjustDisplayMode during initialization
  adjustDisplayMode();

  // Also, call adjustDisplayMode on window resize if needed
  window.addEventListener('resize', adjustDisplayMode);
</script>
```

### Guidelines

- **Responsive Design**: Ensure that your artwork adapts to different screen sizes and aspect ratios.
- **Initialization**: Adjust the display mode when the artwork initializes, before rendering any frames.
- **Testing**: Test your artwork in both modes (`crop` and `fit`) on various devices (desktop, mobile, TV) to verify that it displays correctly.
- **Aspect Ratios**: Be mindful of how your artwork's aspect ratio interacts with the viewport's aspect ratio in both modes.
- **Scaling and Positioning**: Use appropriate scaling and positioning techniques (e.g., CSS transforms, canvas scaling) to achieve the desired visual effect.

## Resize Event Handling

The artwork should handle window resize events properly. Resize events occur when displaying artwork on phones or TVs, especially when changing orientation or when the user adjusts the window size. If the window size changes during artwork display, ensure it triggers a redraw or re-layout with the new canvas size.

Since artworks are displayed in an iframe, resizing the parent window will trigger the resize event on your artwork.

### Implementation

Add an event listener for the `resize` event and adjust your artwork accordingly.

```javascript
<script>
  function onWindowResize() {
    // Get the new window dimensions
    const width = window.innerWidth;
    const height = window.innerHeight;

    // Adjust your canvas or rendering context
    const canvas = document.getElementById('myCanvas');
    canvas.width = width;
    canvas.height = height;

    // Recalculate any dimensions or positions based on the new size
    adjustDisplayMode(); // Re-apply display mode adjustments if necessary

    // Redraw or re-render your artwork
    drawArtwork();
  }

  window.addEventListener('resize', onWindowResize);

  // Initial setup
  function init() {
    onWindowResize(); // Adjust to initial size
    // Other initialization code
  }

  window.onload = init;
</script>
```

### Guidelines

- **Efficient Redrawing**: Optimize your resize handling to avoid performance issues. For example, use a debounce function if the resize event triggers too frequently.
- **Consistent Experience**: Ensure that resizing does not disrupt the user's experience or the artwork's state unless intended.
- **Testing**: Test your artwork's resize behavior on different devices and browsers to ensure compatibility.
- **Canvas Scaling**: If using canvas, remember to update both the element's width and height attributes as well as the style to ensure proper scaling.

Reference: [MDN - Window Resize Event](https://developer.mozilla.org/en-US/docs/Web/API/Window/resize_event).

## Reference

### Source Code

#### IPFS Links

- [V1 - ipfs://QmTBGnS8GdZ1LFFgJzZoM4ArMmxwsrTnEmtEcgLH1h2d1w](https://ipfs.bitmark.com/ipfs/QmTBGnS8GdZ1LFFgJzZoM4ArMmxwsrTnEmtEcgLH1h2d1w)
- [V2 - ipfs://QmcrNv6RihePaRSX7ce5vg9ubtzdz6hjWicqB97uJhGqRk](https://ipfs.bitmark.com/ipfs/QmcrNv6RihePaRSX7ce5vg9ubtzdz6hjWicqB97uJhGqRk)
- [V2 staging - ipfs://QmavGeGEmnfuYPNY1YhNaygj1zN8Z3wk1HrEAnXdu5wSmt](https://ipfs.bitmark.com/ipfs/QmavGeGEmnfuYPNY1YhNaygj1zN8Z3wk1HrEAnXdu5wSmt)
- [V3 - ipfs://QmakVUXKyHF5cddi6owHLbCib7wbU36nY2JoQiNkE7eSho](https://ipfs.bitmark.com/ipfs/QmakVUXKyHF5cddi6owHLbCib7wbU36nY2JoQiNkE7eSho)

### Additional Notes

- **External Dependencies**: When creating your artwork, ensure that all external dependencies are properly included and accessible.
- **Performance**: Be cautious with performance on devices with limited resources, such as mobile phones and TVs.
- **Animation Frame**: Consider using `requestAnimationFrame` for animations to optimize performance.
- **Testing**: Always test your artwork thoroughly in the environments where it will be displayed (desktop browsers, mobile browsers, mobile apps, TV apps).