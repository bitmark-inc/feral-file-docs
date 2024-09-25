# Feral File Artwork Library

###### tags: `Feral File` `operation` 


## Feral File Variables

For Feral File generative artworks, it **injects** pre-defined variables then the collector click to view the software artwork in in frame. This helps artists be able to make variants to their artworks. The variables are:
- `blockchain` - (bitmark | ethereum | tezos)
- `contract` - contract address if any. empty for bitmark
- `token_id` - a unique id of a token in the blockchain. It is hex in bitmark and decimal in tezos and ethereum
- `token_id_hash` - the sha256 hash of a token ID. 

### Token ID

By the design of token contract, it requires a unique id for each tokens. However, in our exhibition structure, we have multiple series and each series would have its own artwork with a same index starts from 0. This brings duplication of token IDs to a contract if we simply use the index as token ID. To avoid confliction, we create a function to generate deterministic unique token id by combining series id and token indexes.

### Token ID Hash

It is the sha256 hash of a token ID. The value of the hash is 32 bytes and represented in hex. Here is an example of the hash
```
0x0aaa657eb7727d08721d447e53ae2959a5a4c6e4a061cc21724a6f759d025d4b
```
which is a 64-length hex string and is prefixed with `0x`.

## Deterministic Randomness

To provide randomness to generative artworks, we create a snippet to offer a random function based on sfc32. The function takes `token_id_hash` (defined above) as the seed of the randomness. This ensures that the randomness is deterministic. If the `token_id_hash` is not given, the snippet will create a random hash so we can test loacally. 

:::success
We use sfc32 in our random function. sfc stands for "Small Fast Counter". It runs very fast in JS. It is part of [PractRand](http://pracrand.sourceforge.net) random number testing suite.
:::

### Code

```htmlmixed
<script id="ffhash-snippet">
  function randomHash() {
    return '0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff'.replace(/[f]/g, function (c) {
      let r = Math.random() * 16 | 0;
      return r.toString(16);
    });
  }

  const urlSearchParams = new URLSearchParams(window.location.search);
  const params = Object.fromEntries(urlSearchParams.entries());
  let hash = params.token_id_hash || randomHash()
  console.log(hash)

  {
    let sfc32 = function (uint128Hex) {
      let a = parseInt(uint128Hex.substr(0, 8), 16);
      let b = parseInt(uint128Hex.substr(8, 8), 16);
      let c = parseInt(uint128Hex.substr(16, 8), 16);
      let d = parseInt(uint128Hex.substr(24, 8), 16);
      let rand = function () {
        a |= 0;
        b |= 0;
        c |= 0;
        d |= 0;
        let t = (((a + b) | 0) + d) | 0;
        d = (d + 1) | 0;
        a = b ^ (b >>> 9);
        b = (c + (c << 3)) | 0;
        c = (c << 21) | (c >>> 11);
        c = (c + t) | 0;
        return (t >>> 0) / 4294967296;
      };
      for (let i = 0; i < 1e6; i += 2) {
        rand();
      }
      return rand
    };

    var ffrand = sfc32(hash.substr(34, 32))
  }
</script>
```

### Usage

After you load the snippet, simply call `ffrand` to get a random number.

```htmlmixed
<script>
  let r = ffrand()
</script>
```

:::warning
If you would like to design the randomness by yourself, please ensure the random value is deterministic by a given `token_id_hash`. Otherwise, the artwork would be different by each page reloads.
:::

## Artwork Library

With this parameters from FF, we come up with a library that can get provenance from our indexer. Thus, for software artwork developer's, they can create artwork variants by integrating this library.

Learning from fxhash, this library is designed to be a simple snippet script (library) that can be put on the top of generative artworks. 

In this snippet script, we only use XHR(XMLHttpRequest). It is because the library does not require any other third-party libraries. This minimize the dependencies of artworks for artists.


### Function

Here lists supported functions

- Get blockchain height (integer)
- Get provenance (array)

### Usage

Attach the script in HEAD. And call to `ffinit` on page loaded.

```htmlembedded
<!DOCTYPE html>
<html lang="en">
<head>
  <script id="ff-library" 
  src="https://ipfs.bitmark.com/ipfs/QmRfJutL1FyAKqT2vAAPErgcJqTfCyWvQtnbiShZhTkNL4"
  type="text/javascript"></script>
</head>

<body onload="ffinit()"></body>

</html>
```

In your script, you are able to get these events:
- provenance-request-error
- provenance-ready
- blockchain-info-request-error
- blockchain-info-ready
    
Example:

```javascript!
 <script>
    // listen to blockchain information data
    window.addEventListener("blockchain-info-ready", function (event) {
      console.log("current block height:", event.detail.height);
    })

    window.addEventListener("blockchain-info-request-error", function (event) {
      console.log(event.detail.error)
    })

    // listen to token provenance data
    window.addEventListener("provenance-ready", function (event) {
      console.log("provenance length:", event.detail.provenances.length)
    })

    window.addEventListener("provenance-request-error", function (event) {
      console.log("fail to get provenance:", event.detail.error)
    })
  </script>
```

### Provenance Data Example

```json!
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
To provide function tracking artwork attributes, we create a snippet to offer a `$feralfile` namespace and 2 functions `$feralfile.features(array_of_attributes)` to input artwork attributes and `$feralfile.getFeatures()` to get saved attributes based on artwork information

### Code

```htmlmixed
<script id="feralfile-attributes">
  var $feralfile = {
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
After you load the snippet, simply call `$feralfile.features()` to set attributes.

```htmlmixed
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
        {
          name: 'color',
          value: 'dark',
        },
        {
          name: 'tmax',
          value: 254,
        },
        {
          name: 'tmin',
          value: 189,
        },
        {
          name: 'letter',
          value: 5,
        },
        {
          name: 'dotnum',
          value: 374,
        },
      ]);
    </script>    
</body>

</html>
```

### TODO

- Refactor API calls
- Merge events
- Select to enable watching

## Reference

### Source Code

#### IPFS Link

- [V1 - ipfs://QmTBGnS8GdZ1LFFgJzZoM4ArMmxwsrTnEmtEcgLH1h2d1w](https://ipfs.bitmark.com/ipfs/QmTBGnS8GdZ1LFFgJzZoM4ArMmxwsrTnEmtEcgLH1h2d1w)
- [V2 - ipfs://QmcrNv6RihePaRSX7ce5vg9ubtzdz6hjWicqB97uJhGqRk](https://ipfs.bitmark.com/ipfs/QmcrNv6RihePaRSX7ce5vg9ubtzdz6hjWicqB97uJhGqRk)
- [V2 staging - ipfs://QmavGeGEmnfuYPNY1YhNaygj1zN8Z3wk1HrEAnXdu5wSmt](https://ipfs.bitmark.com/ipfs/QmavGeGEmnfuYPNY1YhNaygj1zN8Z3wk1HrEAnXdu5wSmt)
- [V3 - ipfs://QmakVUXKyHF5cddi6owHLbCib7wbU36nY2JoQiNkE7eSho](https://ipfs.bitmark.com/ipfs/QmakVUXKyHF5cddi6owHLbCib7wbU36nY2JoQiNkE7eSho)