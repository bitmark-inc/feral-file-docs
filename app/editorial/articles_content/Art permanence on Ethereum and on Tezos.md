# Art permanence on Ethereum and on Tezos

This article was originally published on [fxhash](https://www.fxhash.xyz/article/art-permanence-on-ethereum-and-on-tezos) on February 10, 2023, by [Claus Wilke](https://twitter.com/clauswilke).

An extended article based on a recent Twitter thread on "chain risk" and art NFTs on Ethereum and Tezos.

![](https://i.imgur.com/BqO34eK.png)


There are many big NFT collectors that have extensive and expensive art collections on Ethereum and that refuse to collect on Tezos. Even for artists they like and collect, they will only buy art that was released on Ethereum. This has lead to the well known Ethereum premium or Tezos discount, where art by the same artist is generally more expensive on Ethereum than on Tezos, sometimes by a lot. The argument that is generally brought up against collecting on Tezos is one of "chain risk," which goes as follows: We don't know if Tezos will still be around in ten years, but we are quite certain Ethereum will. So only collecting on Ethereum is safe.

I feel the entire chain risk argument is somewhat silly. It's almost like saying "I don't collect art painted on paper. It could rot or fade or tear. I only collect metal or marble sculptures, where I'm certain they'll be fine in 50 years." If you actually care about the art then you collect the art regardless of medium and you do your best to preserve it, rather than taking a fully passive approach where you expect the art to just always be there without you having to do anything. Tezos is a decentralized blockchain and the mere act of collecting art on it helps it stick around for longer. And, anybody can run Tezos validators, and as long as even a handful of art enthusiasts continue to run validators Tezos won't go away.

Another argument that frequently is brought up is that art on Ethereum is fully on chain (stored on the blockchain itself), and on Tezos it is not. This is a complex topic but most importantly, the differences are subtle and platform-specific and not tied to the blockchain. First, most art on Ethereum is not on chain. NFTs minted on Foundation, or SuperRare, or OpenSea will generally not store the art on chain, but (in the best case scenario) on IPFS (more on this later). In the worst case scenario, art is just stored on private servers (looking at you OpenSea). The risk of losing that art is much much much much higher than the risk of losing the Tezos blockchain. Anybody who feels comfortable buying an Ethereum NFT with art stored on a private server should also be comfortable with buying a Tezos NFT, in particular with art stored on IPFS (which is an almost universal standard on Tezos).

Now back to on-chain art. There is some on Ethereum. Most notably, the generative art released on ArtBlocks has the generative JavaScript code stored on the Ethereum blockchain itself. So as long as Ethereum exists, you can extract the code, run it, and recover the art. But there are some caveats. First, ArtBlocks allows (some) external dependencies, so the work is not fully self-contained. You will generally need some external library such as p5.js which you hopefully can find somewhere. (Some versions of it have been stored on chain also.) Second, you can't serve the art from the blockchain, the blockchain is only a backup mechanism. If all you have a browser and the Ethereum blockchain you can't display an ArtBlocks piece. And in practice, ArtBlocks art is served from a private server. (However, the newer platform [256 ART](https://256art.com/) has solved this issue. Art on that platform can be recovered directly from the blockchain.)

Now let's talk about IPFS (the InterPlanetary File System). IPFS is a great tool to store larger files that don't fit into a blockchain. One nice feature is that you can just serve the art from IPFS. Just stick the IPFS link into a browser that supports it (such as Brave) and there you go. Also, IPFS links are uniquely tied to the file contents. If you have a backup of the files you can always restore IPFS files if they got lost somehow. (That's the whole point of [ClubNFT:](https://www.clubnft.com/) make backup-copies of your own NFTs so you can never lose them.) One problem with IPFS is it doesn't ensure the files actually exist. You could have an NFT on IPFS today and tomorrow it's gone. It's generally up to the NFT platforms and marketplaces to ensure files are available, but they could go out of business and files could disappear.

There are other types of storage solutions that in theory provide persistent long-term storage, paid for up-front. The most widely known one is called Arweave. I haven't looked into it in detail so can't comment much on it, but I have some concerns. What's the correct cost today to have a file stored for the next 100 years? Arweave makes the assumption that the future cost of storage will rapidly go to zero, but what if that's not true? To me, Arweave sounds a bit like a pension system. You pay in when you're in your twenties and you hope there will be funds for you to withdraw when you're in your eighties, and while this is a good idea and it often works it also frequently doesn't.

Let's get back to Ethereum vs. Tezos and common storage modes. On Ethereum, most art (if you're lucky) is stored on IPFS, and some (ArtBlocks, 256 ART) is on chain. On Tezos, almost all art is stored on IPFS. You could store art on chain on Tezos, but almost nobody does currently. fx(hash), the primary generative art platform on Tezos, stores everything on IPFS. However, I would expect that they will eventually provide an on-chain solution also. Again, there is no technical reason why they couldn't, it's just a matter of development priorities.

I should also note that even on-chain art is not guaranteed to never disappear. Blockchains are not in the business of providing storage. If you look at the Ethereum roadmap, much of it is about purging data that is no longer needed. They may very well at some point start removing some of the data blobs that contain art, and at that point the storage guarantees for on-chain art are not any different than for IPFS: You have to make your own backup if you want to be sure somebody still retains a copy.

So all this goes back to: How much work do you as collector have to do to ensure the art you collect remains available. And you probably have to do some, because none of the solutions are 100% automated and self-sustaining. If you have any NFTs that are stored on IPFS, you should really back them up, for example using the ClubNFT tool. If you have any NFTs stored on private servers, it's unclear how to back them up, but you could right-click-and-save (yes, I said it). Also, if you want to limit yourself to collecting only 100% on-chain art you are excluding tons of interesting work, because on-chain storage is too expensive to include high-resolution images or videos or sound samples. On-chain is not the solution for data permanence. None of this is really any different between Ethereum and Tezos, and you'll have to put in some work preserving your art, just like you would if you had a valuable aquarelle or oil painting. And given all this complexity of different storage approaches and failure modes, it sounds like a poor excuse to me to bring up chain risk but not other data permanence risks.

In fact, for generative art, a much bigger risk is that code simply stops working. The saved code on Ethereum won't do you any good if it doesn't run any longer in a modern browser, or if it suddenly produces different output. With generative art, that's a much bigger risk than losing the code in the first place, and people rarely talk about it. It's what keeps me personally up at night. I'm old enough to have seen almost every piece of code I've written break due to external changes beyond my control. It will happen.

Fortunately, there are ways of dealing with this also. As one example, Matt Ebb released a piece on fx(hash) that no longer works (Wandering Nets), and then he released an update which he made available to all of the owners of the original work. We'll see more of this in the future.

![Wandering Nets Redux #1](https://media.fxhash.xyz/w_1400/QmV3cDAk9Ygh13wHW9E62agxf7N3iEHCJxHcQvJbo47eq5)<caption>"[Wandering Nets Redux #1](https://www.fxhash.xyz/gentk/slug/wandering-nets-redux-1)" by [Matt Ebb](https://www.fxhash.xyz/u/Matt%20Ebb)</caption>

One final point I'd like to discuss is the resolution of the art that is being stored. Some people claim that art on Ethereum is more future-proof because it is higher resolution. See for example [this interview.](https://www.youtube.com/watch?v=hiLQNPoiLAw&t=345s)

This is simply false. In particular on fx(hash), where artists deposit the code that generates the art, it is generally possible to recreate the art at very high or even infinite resolution. There are three common strategies that artists employ to make their art future-proof. While some artworks released on fx(hash) may not employ any of these strategies most do. I will describe these strategies in the context of some of my own works, where I know exactly how they work under the hood. I will also link to some works by others that I believe are implemented similarly.

First, some code on fx(hash) generates SVG files, which are resolution independent. Many of my own collections to date are made in this way. See for example "Before/After".

![Before/After #111](https://media.fxhash.xyz/w_1400/QmV4SrZknZQ5dRe4h1EghVJ7874UA3KL61fBhKRha2MH5L)<caption>"[Before/After #111](https://www.fxhash.xyz/gentk/slug/beforeafter-111)" by [Claus Wilke](https://www.fxhash.xyz/u/clauswilke)</caption>

"Hypergiraffe" by Piter Pasma also works this way.

![Hypergiraffe #158](https://media.fxhash.xyz/w_1400/QmVUqLknqspo4WoWoR8RfaEyGuSDfEV9gXzcuM7qBNqqGX)<caption>"[Hypergiraffe #158](https://www.fxhash.xyz/gentk/slug/hypergiraffe-158)" by [Piter Pasma](https://www.fxhash.xyz/u/Piter%20Pasma)</caption>

Next, artworks that run shader code tend to run at whatever resolution your screen is. The shader program is executed for every pixel on the screen. My collection *Voxelsne* is implemented in this manner. And, for *Voxelsne*, if your GPU is sufficiently powerful you can increase the amount of antialiasing to get extra smooth renders.

![Voxelsne #1](https://media.fxhash.xyz/w_1400/QmSMi4D6FfFUgsNt6EyK3RFSz5m67B6qxmFtsSjmGgT9gL)<caption>["Voxelsne #1"](https://www.fxhash.xyz/gentk/slug/voxelsne-1) by [Claus Wilke](https://www.fxhash.xyz/u/clauswilke)</caption>

A famous collection that similarly adapts to the screen size is *Reading a book* by Kim Asendorf.

![Reading a book #831](https://media.fxhash.xyz/w_1400/QmQhVHJTsVELt6aS3tqDzVb9Bu3Bp9xUUvxqoo3n7PmEZb)<caption>"[Reading a book #831](https://www.fxhash.xyz/gentk/slug/reading-a-book-831)" by [Kim Asendorf](https://www.fxhash.xyz/u/Kim%20Asendorf)</caption>


Finally, probably the most common approach is to render into an off-screen canvas at a given resolution, and to allow the users to request a higher resolution via keyboard commands or URL parameters. My collection *The Artifact* works in this manner. You can try out [this link](https://gateway.fxhash2.xyz/ipfs/QmVSse7Y1Vw3NbwCbjuvtPn7rqL1bvXvBDkTsiL1RVVDeQ/?fxhash=oo5BehjnhSgz6Mt3e5PQ1zCLW9winX71bBMZ7WfzguVaYi1HXCD&r=2&d=2) to generate an ultra-high resolution render of the same iteration shown here.

![The Artifact #220](https://media.fxhash.xyz/w_1400/QmXj4cAczDtuKF1CDSFAKxMBdF2ucPbywKaXxUeNyjJFbV)<caption>"[The Artifact #220](https://www.fxhash.xyz/gentk/slug/the-artifact-220)" by [Claus Wilke](https://www.fxhash.xyz/u/clauswilke)</caption>

The collection *Stepping Stones* by teaboswell has similar functionality.

![Stepping Stones #757](https://media.fxhash.xyz/w_1400/QmfHGXZcJLhWYAbdWtQarbJbjWtuo1gQDob4GMSfDjzd2u)<caption>"[Stepping Stones #757](https://www.fxhash.xyz/gentk/slug/stepping-stones-757)" by [Tyler Boswell](https://www.fxhash.xyz/u/teaboswell)</caption>

In conclusion, regardless of the blockchain used, there are concerns about data permanence and reproducibility of the art in the future, though most of these concerns can be addressed with reasonable effort. But collectors will have to play their part in preserving the art. In addition, it would be great if there were a non-profit foundation whose purpose it is to provide backups for all art-related IPFS assets, similar to what the internet archive does for web2. We shouldn't all individually have to worry about permanence of our collection, even if we make backups. Ideally this foundation would take donations on-chain so we could just send some portion of our primary or secondary royalties to them whenever we mint NFTs.

P.S.: This article is based in part on a [Tweet thread posted on Feb. 5, 2023.](https://twitter.com/ClausWilke/status/1622401965337595908)

P.P.S.: The creator of the Ordinals NFT platform on Bitcoin recently brought up many similar points [in this Bankless interview](https://www.youtube.com/watch?v=ktL77zEWcEc&t=1451s), plus some additional ones that also apply equally to Ethereum and Tezos. It's worth a listen.
