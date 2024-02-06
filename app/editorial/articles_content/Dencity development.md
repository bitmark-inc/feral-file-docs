# Dencity development

This article was originally published on [fxhash](https://www.fxhash.xyz/article/dencity-development) on September 1, 2022, by [Hevey](https://www.fxhash.xyz/u/Hevey).

### The development of the [Dencity project](https://www.fxhash.xyz/generative/slug/dencity) was particularly complex and required months of intensive work... Here is an EXTREMELY SIMPLIFIED summary of the many stages of the development of this crazy project.

![](https://i.imgur.com/ADoRA6i.png)

**Connections**

The first step was to create a grid and randomly place points on it:

![](https://gateway.fxhash2.xyz/ipfs/QmZkewjrLuZY5DD2dnafFh8kEpuqwZyXf4TzpFGjtcbzVg)

These points were then joined together horizontally:

![](https://gateway.fxhash2.xyz/ipfs/QmajuZ5VTZ3PPAKVXk9qjBMxVHjEEdAMx1zVshbE4S2tkn)

The rule was as follows "connect to the first available point, from left to right"... To ensure that the result was correct, blue triangles were drawn temporarily on the links to visualize them well:

![](https://gateway.fxhash2.xyz/ipfs/QmezaEXsVA1hgmFWjdakng1AaWNGw56WFQJG7tPiCXc9h2)

Then the vertical links were added:

![](https://gateway.fxhash2.xyz/ipfs/Qmb3Mi1HaSBbhVx2aR9qD9Y4t4DLKX5ywZXUdQC6TC2TUS)

**Perspective effect**

Before continuing with the roads, I created the function that calculates the perspective effect and tested it with rectangles:

![](https://gateway.fxhash2.xyz/ipfs/QmfRyLikLUWRoyryx2c3NZJXm6VfJbb7CfT9Ay73KiAnar)

It was then enough to use the calculated positions to create a block with a 3D effect:

![](https://gateway.fxhash2.xyz/ipfs/QmdJwfNDjXzmgLmzNjaitGpYiFFcPkVjJScSYHbcWiSaxL)

**Grid**

I then took my code with the links and adapted it to the 2/3 format:

![](https://gateway.fxhash2.xyz/ipfs/QmZCuddQR5ENai4wKTXyx2PPwwT1eqfMncZQ59cL95gWBm)

Broad lines have been drawn:

![](https://gateway.fxhash2.xyz/ipfs/QmW9yn5rBu6QZrjAW1XGNT93QBrwUrHPbmEeWkvffXcj8m)

White lines have been added in the middle of the roads:

![](https://gateway.fxhash2.xyz/ipfs/QmUU1WAee9u8ErRqYMNFRsiQrzkd5VPAMRqtgM2x8tLN56)

The white lines were then removed at the intersections:

![](https://gateway.fxhash2.xyz/ipfs/QmbTxrM2h5jbnSp2XfoE6JNxXCdF53VVrbvAbpaX84e6N7)

And pedestrian crossings have been added near intersections:

![](https://gateway.fxhash2.xyz/ipfs/QmShWUUcYLncEzdCru63q3K3ZLkrjtevqQhiAZLgm5ew4d)

This was, however, only a draft version of the roads which were completely redesigned later in development.

**Placement of buildings**

To be able to visualize the grid (technically contained in arrays), I created a function that draws the grid over the result.

The red "r" correspond to the squares occupied by roads, the blue "t" are the sidewalks around the road (this allows to leave a margin so as not to put buildings too close to the road) and the gray "o" are free squares:

![](https://gateway.fxhash2.xyz/ipfs/QmZJ5TwhH4GfLctAqavXhsuLeNViBCc26S5g6n9zZ9fneg)

It was then necessary to randomly place areas of variable size in the free squares (the green "i" squares are those of the buildings):

![](https://gateway.fxhash2.xyz/ipfs/QmaA2u4KT13GqxetEg6SDeRayDikhbYxNXn2M7XgeAT6Bb)

From there, I took the piece of code created a little earlier and placed the first version of the buildings on the green spaces:

![](https://gateway.fxhash2.xyz/ipfs/QmVDEKbBpZQVPHRRjEcupFHXjtbdJM7UEdcAT7kgiiYs8C)

We also had to create a complex sorting function to classify the buildings according to the order in which they must be drawn, otherwise we obtain this kind of result:

![](https://gateway.fxhash2.xyz/ipfs/QmYLJyxoFBj7Ywnofsn7iTyu7M3YetDdcFbqG2ULbR4Xdb)

Then by expanding the grid, hiding the debug grid and placing as many buildings as possible, we get this:

![](https://gateway.fxhash2.xyz/ipfs/QmTuTVv9JLzQevBiFNirW599ZWBTUbx76H8MBEJgoubBnc)

**Building facades**

I started to draw on the facades of buildings:

![](https://gateway.fxhash2.xyz/ipfs/QmbzsVhcsKE76nCdo89GcM6u2fWD1Cu1oBLCb5JhFvLxZt)

And to optimize the execution speed of the code, I tried to determine which facades are useful to draw and which are useless:

![](https://gateway.fxhash2.xyz/ipfs/QmXiqgJuheLKJhHvAUAhsriRUP9xkgxi8tWMsRfX3bWTBw)

The first type of facades was the one with small windows:

![](https://gateway.fxhash2.xyz/ipfs/QmaD9NpDA1mWzBV26GEoCzEEq2ygwSvsrsti7atnMMvamj)

The size of the windows was then randomly defined for each building:

![](https://gateway.fxhash2.xyz/ipfs/QmTJ9byH7DuAybbPBTsPz7BeJ3iBCJCjCKfni6jDRY4wDL)

Then a second type of facade was created:

![](https://gateway.fxhash2.xyz/ipfs/QmXRymqUUL3TKXC693zuwvZuwChDqWA2CvxWKKa7Gpjxc1)

And the type of facade was randomly defined for each building:

![](https://gateway.fxhash2.xyz/ipfs/QmfC9zVbuU8N8u3hLoiVKwroZGBDaXaP212CdUVBAfbmwj)

**Roof edges**

The next step was to create the roof edges by first placing shapes (similar to buildings without windows) on top of the buildings:

![](https://gateway.fxhash2.xyz/ipfs/QmcsFmDaAw45T9GFbyz3hUmYhv3tfjKTRtdE3sNDSPh3Fj)

The addition of these roof edges depends on the type of facades of the building, only the first type of facade is entitled to roof edges at this stage of development:

![](https://gateway.fxhash2.xyz/ipfs/QmWR9EL9UHMEAhVN9EMEUHm4FSyPcwKaN7oawR4diX91ZF)

**Drawing effect**

To begin with, the borders were drawn around the different elements:

![](https://gateway.fxhash2.xyz/ipfs/QmcFoiufFbh8Hme6zEbzpDRk3EQbHSoLNi9RKughwDgCnS)

The different facades were then modified to create a small random shift in each line and obtain a hand-drawn effect:

![](https://gateway.fxhash2.xyz/ipfs/QmUyfukWmnjPSugQgFz7NUnwZhGYijjEiCcW6ywmmvX3Dz)

The lines of the buildings have also been modified to obtain a slight offset and create this drawing effect:

![](https://gateway.fxhash2.xyz/ipfs/QmSt1npTPzQLbmtvDWLTLetDunksPrgUDFguSVCD2dXm2L)

Then some edge lines were removed to give the illusion that these edges are one with the building:

![](https://gateway.fxhash2.xyz/ipfs/Qmb41sn5zsozdZ31ZDmFj7BqMPntKx8KhzSzEnHuQvPVCU)

**Round building**

A new type of building has been added to the project, the round building:

![](https://gateway.fxhash2.xyz/ipfs/QmTVCiu2VALewcpbAHQYLu82tbfsgCmQdKGezXeZrU4JVE)

Each type of building has types of facades defined for it... The round building can use the second type of facade created but cannot use the one with the small windows, on the other hand it is entitled to a new type of facade (the one at the top left):

![](https://gateway.fxhash2.xyz/ipfs/QmewrpcBh88wYBsCFeSGKH5fH3aEGwHQF1AVh3yRVtaWZF)

**Sections**

Some buildings are sometimes composed of several sections. From a technical point of view, it corresponds to a building on top of another building:

![](https://gateway.fxhash2.xyz/ipfs/QmSFGd1oYvpRubvBJ7XmhP8oBEZuXLjr2RZ9k7rmvmPtid)

To create a layered effect, simply reduce the dimensions of the sections:

![](https://gateway.fxhash2.xyz/ipfs/QmRVDiFb8qQ7hiQJFWeRWHxUZdgRFsEFtoW2xgSnwUJbau)

**Roof accessories**

To add even more detail, the first roof accessories have been randomly added to certain buildings:

![](https://gateway.fxhash2.xyz/ipfs/QmYBLCVDacTnDhiSTmf5UUCpiXJTbYTjBQapDUQvP8bioc)

With sometimes a second element on top:

![](https://gateway.fxhash2.xyz/ipfs/QmbAGTzbMinUB997knZ42mpPTQLYrDJXJBFWiE4mRUjrPR)

You can also see that a new type of facade has been added.

**Redraw the roads**

In order for the roads to have a drawing effect similar to the buildings, they had to be completely redone.

The function that manages the perspective effect can be adjusted to have a greater or lesser effect. In this case, the effect has been reduced to better see the roads:

![](https://gateway.fxhash2.xyz/ipfs/QmbBLEmjg18pjuALmX4sGk81Pp5YxgT6NHxNqi57wRNec8)

Irregular borders have been drawn along the roads:

![](https://gateway.fxhash2.xyz/ipfs/QmNr5bHvqhA7ordPGmwn9EVfygP9ybdNQDW4QvLMPPtBjY)

It then took a lot of work to ensure that all the lines were well drawn everywhere except at the intersections and that the lines were still correctly connected despite the random offsets (necessary for the drawing effect):

![](https://gateway.fxhash2.xyz/ipfs/QmTN3z8dJEizyBGdZmaUW7p2nEZ9izP4VEHjFJuEsGK32P)

Then the white lines and pedestrian crossings were added:

![](https://gateway.fxhash2.xyz/ipfs/QmaXqFcGyZi27m94uofocjpR8117ZxftdoGGQLAJMQ3s5J)

**Colors**

At this stage of development, it was too early to define the final color palettes, so I retrieved the palettes from a previous project to be able to immediately do the first tests with the colors:

![](https://gateway.fxhash2.xyz/ipfs/QmVVKsNWcXLLcCEFL78RNL6okHrWZnvx7ttGEbYDiY37Mj)

Just like the perspective effect, the function that creates the offsets for the drawing effect can be modified to create a more or less pronounced effect:

![](https://gateway.fxhash2.xyz/ipfs/QmRYgKcFe6pGPY2E8BqG7RsEM4ypnZVZHAaRBkA1fxfx8S)

The color of the roof edges was then lightened and for more details, blinds were also added randomly:

![](https://gateway.fxhash2.xyz/ipfs/QmUjk9yP4tx3FkHmbg12eJXtWKSvDz9bSJF5apDWX6ZWoS)

**Hexagonal building**

A new building type has been created:

![](https://gateway.fxhash2.xyz/ipfs/QmdMUyTPXDgd6Ls6apqhotzz2CtAh4Gtv2n1aobud9Vs28)

And bugs related to diagonal facades have been fixed:

![](https://gateway.fxhash2.xyz/ipfs/QmUHLeX6H8QSxxtuaPiBM8iRQ7eUcrABuNLJ3TytXvUfbj)

**Roof accessories**

To bring a little more variety to the roof accessories, on some buildings they will be added several times (and to better visualize the lines, the drawing effect has been temporarily disabled):

![](https://gateway.fxhash2.xyz/ipfs/QmRmTwCcPJwnbggUybZngXfjVbdohMm7wi3hRJt8hNwWMp)

These accessories will sometimes be added by 2 or more rarely by 4 on the roofs:

![](https://gateway.fxhash2.xyz/ipfs/QmVhG9Lfm7Q1TKJ89nKqB18oBwABhdtpKHtaRQg1FfGat4)

A new accessory for round buildings has been added, however it was not centered:

![](https://gateway.fxhash2.xyz/ipfs/QmU8rXC77i4BkBEFBdWsxj3ubpn7NDSTuj2ZgpZJjhAgMp)

To solve this kind of problem, the debug tool was used very often during development... We can see here that the accessory was well centered and that it was actually the building that did not occupy the whole area:

![](https://gateway.fxhash2.xyz/ipfs/QmcedQjm43vu37e1V7NLgzuMK6dXfHcST7PMtGgCiSUHrZ)

It was therefore necessary to modify the dimensions of the building to fix this problem:

![](https://gateway.fxhash2.xyz/ipfs/QmVk6MuKiLbPTP3xRc3VK6oVoBDtVLfz7ZqMtcaM5nQWjQ)

Hexagonal buildings were also entitled to their accessory:

![](https://gateway.fxhash2.xyz/ipfs/QmYVhuXDMEWTpNo2ekMSyTh5GUSR1Jn9tBL4vDG6k5huBm)

**More complex buildings**

To further complicate the project and add even more details, I then created a new type of building with many more facades than the rectangular buildings:

![](https://gateway.fxhash2.xyz/ipfs/QmRSkZnA9QjztxEs82AEsGybyx1XTFmLC9AdaXPdzHbjnu)

A second version with some of the diagonal facades has even been created:

![](https://gateway.fxhash2.xyz/ipfs/Qmcj8ViwywyoH5MsUqPHq7dPX72VGG95snqKow1VmLF52E)

It was then necessary to correctly manage the facades on these complex buildings, and more complicated still, to add the edges of the roof:

![](https://gateway.fxhash2.xyz/ipfs/QmZgM2RptHVGZeLBkxVrnPPu6dXT7T1YwYV38ENV5Adty4)

After a long test and debugging work, the roof edges were finally displayed correctly:

![](https://gateway.fxhash2.xyz/ipfs/QmVzz7MncjqVUBfZH1uwRChqeRbzPM44g1tV33cr27qCNg)

**Hospital and heliport**

The new buildings have a suitable appearance for a hospital, so I decided that the biggest of these buildings would have a chance of having a helipad:

![](https://gateway.fxhash2.xyz/ipfs/QmQBzCT3FJx9Z9jh7eK3AJkxnzrrxze4qvG6TApHXHAAJP)

With an "H" drawn just with lines so that you can then offset the lines slightly for the drawing effect:

![](https://gateway.fxhash2.xyz/ipfs/QmRXw622wthvA5KiEMQ7P5u1irzxsEVyYjgJxSHdFWviUg)

**Large buildings**

I tested different sizes of buildings to define the limits to apply and avoid obtaining too large buildings:

![](https://gateway.fxhash2.xyz/ipfs/QmUWJZEPF91i1ZhifBMZqZEjkyTo52swEvkqpgvehatofX)

To further vary the result, I then added towers much taller than the equivalent buildings:

![](https://gateway.fxhash2.xyz/ipfs/QmeRDzbU9Qrjdr6BrdeW9UpDYXq2wsZRYQrGpeXsZ6mu4D)

These large towers were entitled to facades just for them:

![](https://gateway.fxhash2.xyz/ipfs/QmSUCfYKk6zAZtzQPKrBX3ghxN5xvsYdFsczkB1xWJWpWL)

**Textures**

A subtle texture effect has been applied to the different elements (instead of a plain color facade):

![](https://gateway.fxhash2.xyz/ipfs/QmZSGBovLJj8anPTg5sNeTQe3VwqbuUU977bMKH9uqAMJT)

This allows for a much more organic result:

![](https://gateway.fxhash2.xyz/ipfs/QmdqfHLLBbaXBP9vypaLH3miD4uCANNw2qacw9cq1DG14H)

The texture effect is created using a large number of lines that vary very slightly in color:

![](https://gateway.fxhash2.xyz/ipfs/QmTAX8XvJPYyKj8bqjsWP4KC6A9VagLmaB6BiYQJm8HrWn)

**Downtown**

Until then the distribution of buildings was random, so for more variety I created the "Downtown" feature which prioritizes the large buildings towards the center and the small ones towards the edges:

![](https://gateway.fxhash2.xyz/ipfs/QmaSpTuxwXm5DeSwQuMJVRSxQLSvtcFsSS9g1boptFDNYZ)

**Diagonal view**

I then started developing another feature to change the orientation:

![](https://gateway.fxhash2.xyz/ipfs/QmU941NJ2HBwXXPXKwE35qDvv7EkDpQDc6MRYCySDPiRfi)

To obtain a 45° angled view:

![](https://gateway.fxhash2.xyz/ipfs/QmQaESKSgUkrcd82qiLzj9ADUQUfGrYJ7qk3npcrRKdEzV)

**Trees**

An important step to add ever more details to this crazy project was that of the trees... To do this, I started by activating the debug mode:

![](https://gateway.fxhash2.xyz/ipfs/Qma4vQGYXURwBjBNBCMibTp66CGdGMDjVZAeBxifWvBA1z)

I then placed "a" in the boxes that were still available:

![](https://gateway.fxhash2.xyz/ipfs/QmZsS46spUJbeksc8BMDqxZfhsEmnFszMCBeAuneiNUyFy)

And I added buildings on these squares:

![](https://gateway.fxhash2.xyz/ipfs/QmW121GkLgEu1bK1NGTF79WfSJhzZtXUX6g3dbFSprMShc)

Why buildings? To optimize and simplify the code... From a technical point of view, buildings and trees have a common Class, this makes it possible to make certain functions accessible to both buildings and trees (without having to rewrite them twice) and it also allows trees to be compatible with the sorting function of buildings (because the trees and the buildings must be sorted together in order to then be able to draw them in a very precise order, otherwise we would see the trees in front of or above the buildings).

The trees are kind of very special buildings... Here you can see that the building is divided into 2 sections:

![](https://gateway.fxhash2.xyz/ipfs/QmernT9Hk25YfMsRwAeErsmXkFyn8i7magBbVbnJefkEfD)

Then a third section was added and the drawings on the facades were removed:

![](https://gateway.fxhash2.xyz/ipfs/QmSGgaXiqjB2onwYFZjpMYrxpckMs2zd7N4t5QQYaS7Qg5)

Other sections have been added:

![](https://gateway.fxhash2.xyz/ipfs/Qma2oXCKMAM7ZRwPvtYdBUzq8aSaTQro35yjXaaRE7W2CJ)

The drawing effect makes sections less visible:

![](https://gateway.fxhash2.xyz/ipfs/QmVwKoPwi6XCtysDST2L2kHQGymWQFGfPtmuVurkHddQKE)

For a less square result, tree sections now have 5 sides, with random rotation of each section:

![](https://gateway.fxhash2.xyz/ipfs/QmNx1QjW2HULn28tuWEjRknCwKeVkVgmz79MUG31oRnUMf)

And to finish, light lighter lines have been drawn on the facades:

![](https://gateway.fxhash2.xyz/ipfs/QmcuKevQnioZzg3Pzc7cXvNpUiBf6wEEHEEJ2NT3WESptV)

**New special buildings**

As if there weren't enough detail yet, I added the pyramid building:

![](https://gateway.fxhash2.xyz/ipfs/QmVXcQUVD6nvW2FhZMP9htUVSPBhvtGiayaPTu8nTAAasz)

Spiral buildings:

![](https://gateway.fxhash2.xyz/ipfs/QmUdrznFdUmZzW5VzREufiKPQKamzKLaDwnRnz3q18nkAH)

As well as many possible variations for spiral buildings:

![](https://gateway.fxhash2.xyz/ipfs/QmcpEQqn9Uy9cHQ8c92ZVdes7zterYuUW6V1QY7euK4CKg)

**Vehicles**

To add the vehicles, I first re-enabled debug mode:

![](https://gateway.fxhash2.xyz/ipfs/QmRsAt6Tz1WEYXsNo9TGAr3PabBbfTvCwRFzwjvkrdfzEx)

I then added "v" boxes on the roads:

![](https://gateway.fxhash2.xyz/ipfs/QmTzC8vKPTUCdNDRnW7aXsbX2AwTAx88capT9v3QafiXZ7)

Just like trees, vehicles are technically a kind of simple rectangular building:

![](https://gateway.fxhash2.xyz/ipfs/QmZ8u99eaxtUBqN74Bww1QRgXQdmKeUSbjefLWfPAWnz4s)

A second section is dropped on top of the vehicles to give them a car or bus appearance:

![](https://gateway.fxhash2.xyz/ipfs/QmSTibZEz13xHV1HjygmC829f5KLS72BDFSR1DMwihb8kB)

Have you ever noticed that the direction of the vehicles has been calculated and that they are all driving on the right?

I then added the vehicles on the other lane in the other direction:

![](https://gateway.fxhash2.xyz/ipfs/Qmcanu7H18jpzMhYJy3oG1TEd2rrihQihK7YxbPajYpep8)

Then I repeated the operation for the horizontal roads by adding the "v":

![](https://gateway.fxhash2.xyz/ipfs/QmXQbbsWYtCtBvAxDryC2xNvpCxAg3szMDQ4kRLsj3GPqG)

To get this:

![](https://gateway.fxhash2.xyz/ipfs/QmNQq2uJhGz5rY72ZaiUeYaxFqqHLxL2Qmr3YhkLKFfRbQ)

**New roof accessory**

A new, much larger roof accessory has been added:

![](https://gateway.fxhash2.xyz/ipfs/QmaTUfLPsJbqS13LZjpeQCmjqpLZT8gi1wdeLY5iN1CS81)

A more or less fine grid was then applied:

![](https://gateway.fxhash2.xyz/ipfs/QmavSt2cr6j6zmo7k2VB6iBQ1YnNaFXYd76TALtgboezzF)

**Parks**

The addition of the parks was also an important step in the project... To begin with, it was necessary to select a part of the areas entirely surrounded by roads, of square or rectangular shape (therefore without a stretch of road inside the area):

![](https://gateway.fxhash2.xyz/ipfs/QmcERXA2gsnCkLWSw3esUMuoasWVaXPh8VEkVrHdcvFmnL)

A lake is then drawn in the middle of the park:

![](https://gateway.fxhash2.xyz/ipfs/QmYeuzBDDUHGWNTTNFwXHJGNw7Yc16NwojygkUvWLTWt5n)

The area occupied by the lake (+ a margin around the lake) is made unavailable for further additions, then the rest of the park is marked with the letter "p":

![](https://gateway.fxhash2.xyz/ipfs/QmWa1PZPSqd8WrNA2hgXAfndWkcty1qrLC6kuPvfPKKqdx)

Trees are then placed in the squares of the park with the letter "p":

![](https://gateway.fxhash2.xyz/ipfs/QmcSzrZW8hWczZCvHUGyUJwfWwWNa7BtJy6qjKfiWw6XeY)

With several parks and the perspective effect, we get:

![](https://gateway.fxhash2.xyz/ipfs/QmYGCKAqCjRJkGQhbxqNfT7BueDtADjTKjTUM5ebyAPSNM)

I then placed other elements in the park resembling picnic blankets to add more detail to the parks:

![](https://gateway.fxhash2.xyz/ipfs/QmY4HrgoCKwo4oHGV8raBzQUE5hALEcPxXoR6VYjJpsBkP)

**Swimming pool**

One of the most original details of this project is certainly the rooftop swimming pool... To do this, I started by adding a section to the building and which occupies a little less than half of the building in the direction of the length:

![](https://gateway.fxhash2.xyz/ipfs/QmQSjera5PmcfF44m5bex5TA4Vzywo4w5asBWjTBFfLxLq)

Since this is a building section, I was able to apply roof edges to create the edges of the pool:

![](https://gateway.fxhash2.xyz/ipfs/Qmbw5RKmHeeKYFneJ25nB2LB943yYxZDzyt7KfXKJ89kAi)

The next step was to add a new element that looks like a bar that serves cold drinks:

![](https://gateway.fxhash2.xyz/ipfs/QmcdH2QWUU2bhRKYPEieBDQnqmZRPY6MpLTpKwaEcgAR1r)

It was then necessary to target the area of the roof on which the deckchairs will be drawn:

![](https://gateway.fxhash2.xyz/ipfs/QmbRhV2jAj8VSxRkZTnsRWUSk6S2TViKXFv9Zr1Nt6rp2r)

Then draw them, sometimes leaving the space empty:

![](https://gateway.fxhash2.xyz/ipfs/QmWG5eXAbuFH6M3zFyXTzGRj8uJde6QYtsxZTjzb4M24vF)

The pool bar is also sometimes drawn in a park:

![](https://gateway.fxhash2.xyz/ipfs/QmRfvc4KEr21vdvmdJtLmEnPZiEHa57SsuKRxf7hfk4YzZ)

**Color palettes**

Although the selection of the palettes had already progressed in parallel with the development of the project, it was at this late stage of the development that I did the most work on the color palettes of the project.

![](https://gateway.fxhash2.xyz/ipfs/QmY8mb9k2Xbj9LT1QBH8BNHRmtNWF4eZFhZDnTge2Hj6MT)

This is also where I created the single color palettes:

![](https://gateway.fxhash2.xyz/ipfs/QmP3RbPV2Hqh2QQYEp8oi4HCWM8TYNVdFumsLoH27wbQ33)

Some of which had an extremely low probability of being minted (about 0.15% chance for this one):

![](https://gateway.fxhash2.xyz/ipfs/QmbEWfxPJLAiACD3R8P9UdN1iyMHTC9Efyg2GGybKFmmvx)

In the end, 77 color palettes were selected:

![](https://gateway.fxhash2.xyz/ipfs/QmUg4y6Dh2sFxaUQdhDpnzrwdVu5pLCccEMHBrAWrmQBdo)

**Low density**

The rare "Low density" feature is rather simple to create, since by limiting the number of buildings, it leaves a lot of room for trees:

![](https://gateway.fxhash2.xyz/ipfs/QmX1KxnAg2ZyXDU4UXgiWNfzE8p9TWbCfKSTviSpVtnF6u)

However, it was necessary to modify or limit many possibilities for this feature... For example, the map could not be too large (because it was much too slow to generate), nor too small because it was unattractive. The number of vehicles has been reduced, some features made unavailable and all palettes have been tested (and only a few palettes have been retained for this feature):

![](https://gateway.fxhash2.xyz/ipfs/QmX5FWwrSm74mUJBM7enhyZvnzibDNjotr4kw2Sa5yPGu4)

**Antennas**

The last roof accessory added was the antennas:

![](https://gateway.fxhash2.xyz/ipfs/QmcXj3wKS6NxFsRApK1nSBQy9fqbYqMaAv1S5DrfvY1qHv)

**Tests and latest improvements**

The last step consisted in carrying out numerous tests and checks, and tweaking all the last little details to obtain an impeccable result:

![](https://gateway.fxhash2.xyz/ipfs/QmbMpmaefCtKeaKcRSbrVBfdRxmiFmnhZsHJKzHaWEv6rL)

**Code size**

The final version of the project code was about 155,000 characters (about 3600 lines of code)... After reducing the size of the code as much as possible using various techniques, the code was minified into a compact code block of 37,934 characters:

![](https://gateway.fxhash2.xyz/ipfs/QmRQK3SH8gDsSbMB5x6775Z2ita4CfR3WsHk8aye2AqwTo)

**About this article**

Of course this article is only a very simplified outline of the huge and very complex work behind the Dencity project.

**Release**

The Dencity project was published on May 4, 2022 on Fxhash and is now available on the [secondary market](https://www.fxhash.xyz/marketplace/generative/12431).

![Dencity #1](https://media.fxhash.xyz/w_1400/QmbzMjCS7FW6mzLS1HfHY3wdotvhR5Wk4fTurm1LuWyvZY)