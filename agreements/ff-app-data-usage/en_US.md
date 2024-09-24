# Feral File App Data Usage

**Last Updated: 20-SEP 2024**

## Why Does Feral File Collect Usage Data?

We collect usage data to improve the Feral File app’s user experience and performance. Since we can’t (and definitely don’t want to) look over anyone’s shoulders to see what you’re doing, the usage data provides an anonymous way for you to contribute to making the Feral File app better.

## How Does Feral File Protect My Privacy While Still Collecting Usage Data?

The Feral File app sets a new standard for protecting the privacy of any usage data you contribute. As far as we know, no other app or crypto wallet does as much to conceal your privacy. If you opt in to sending us usage data, the mobile app will mask your identity and device UUID by hashing them locally on your device before uploading them to our analytics server. The final output of the hashing function is an essentially random string of numbers and letters that cannot be reversed to reveal any information about you, your device, or any of your accounts. The hashes function as anonymous identifiers that help us understand how you’re using the Feral File app without revealing anything about who you are.

Here’s an example of such a hash:

f0c63b1c38cfb46ddca43c02a0b1b29808f3747bc6c7bc399848de7e602b96b1

Anytime you use one of the Feral File app’s main features, such as linking with an external wallet account, your mobile app will anonymously send a message to our analytics server that logs the general type of action and any relevant details. Here’s an example of a log message for linking to an external account:

{  
  "name":"link\_feralfile",  
  "user\_id":"d593d496d6f3dccb8abb51ac7bf627f985a89595a73e037ecc6c80bd",  
  "device\_id":"4a8aabd47458f9a0ecbd0f26f838cd4898a915646f85d4f6b1154bb1",  
  "timestamp":"1652691627636",  
  "platform":"ios",  
  "version":"0.27.0",  
  "internal\_build":**false**,  
  "hashed\_data":{  
    "address":"b9fa7c430edb2e35769ff08ebfff975f5401bdfd7b576967ec4d6905"  
  }  
}

The only reason for hashing your identities is that it helps us identify when the same person is repeatedly trying to accomplish the same action in a short period of time. This is important because it is usually a red flag that we need to fix something. For example, if we notice that you tried to add the same wallet connection repeatedly over the course of a few minutes, it’s a big red flag to us that there’s something that we need to review and fix in the wallet linking process.

## What Kinds of Usage Data Are Contributed?

The Feral File app currently logs usage data for the following actions:

- Creating a wallet account  
- Linking to an external wallet account  
- Submitting the user survey  
- General measures of usage time

## Will Feral File Ever Share This Usage Data With Anyone?

Absolutely and unconditionally not. We will never sell your data to any third parties. We will never allow any third-party advertisers or data collectors to track any of your behaviors in the Feral File app.

## If Feral File’s Servers Get Hacked, Could My Usage Data Be Compromised?

Absolutely not. We will never collect any personally identifiable information about you, such as your name, email address, or geographic location. Your Feral File app subscription information is completely controlled by Apple and Google, and we will never have any access to it. Contributing analytics can in no way put you at personal risk. Since you and only you control the keys to your wallet, no one can ever gain control of your wallet by hacking the Feral File’s servers.

