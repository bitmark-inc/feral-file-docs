# Tips to improve your generative artwork

This article was originally published by [Tyler Hobbs](https://tylerxhobbs.com/) on January 21, 2021.

If you're getting started with generative artwork, I'd like to offer you some advice to help improve your work. I've listed some of the major stumbling blocks I see beginning generative artists make, as well as some discoveries I've made with my own work along the way. There are certainly exceptions to most of my points, but on average, I think you'd do well to try out my suggestions.

**Use HSB for colors instead of RGB**

RGB is the devil. Seriously, it does not match how we think about color AT ALL. This is especially true if you want to "change" colors along the image, using a gradient or some other technique. HSB (which stands for Hue, Saturation, and Brightness) fits our mental model well, and makes those kinds of operations natural. This change alone will improve your color work. I wrote some additional details on using HSB in my article on [working with color in generative art](https://tylerxhobbs.com/essays/2016/working-with-color-in-generative-art).

**Black isn't the only background color you can use**

A strangely high percentage of generative artwork is made with black backgrounds. It's not impossible to make good artwork with a black background, but it can be tricky to do well. That amount of black gives you less room to maneuver when choosing colors. I'd recommend you experiment with other background colors until you're very comfortable working with color. White (or off-white) backgrounds are a good place to start. As a bonus, this will typically look better when printed (see my last point).

**Don't use random colors**

I see lots of beginner generative artwork where RGB colors were selected with random values for R, G, and B. Under almost all circumstances, this is going to look worse than a color palette that is carefully selected. You don't have to be a color genius to avoid this. With generative art it's easy to experiment with different color approaches and fine tune them. I recommend that you start simple: work with one, two, or three distinct colors. Do not select these colors completely randomly. The odds of success are practically zero. Instead, be patient, and take the time to carefully evaluate your color choices.

Here's a demo of the difference this makes. With random color choices:

![](https://hackmd.io/_uploads/BkeBw5wNn.jpg)

With hand-picked colors:

![](https://hackmd.io/_uploads/B10HDqP4h.jpg)

**Try to have 'shapes' at all sizes**

Large, simple compositional shapes provide structure for your image. If you can make these out when looking at a thumbnail of your image, you're doing well. Start with these large shapes, and make them look as good as possible before you focus too much on details.

Details turn your sketch into a finished work. Fortunately, with generative artwork, details don't have to be laborious like they are when painting and drawing, but they do take time and attention. When working on details, I ask myself questions like, "Does that line need to be straight, or would curves look better? Maybe dotted curves?" and "Can I replace that solid polygon fill with a soft texture?" Having good detail is the result of carefully considering every element of your work and asking if it could be better. It's possible to go overboard, but if you're just getting started, odds are you're not doing enough of this.

![](https://hackmd.io/_uploads/HyT0wqPE2.png)

**Experiment with values for every variable**

This is the biggest source of "putting in the time", but it's what will produce the best possible output. It will also allow you to discover ideas you would otherwise miss. Look through your program for any hard-coded numbers. By the time you're done with your program, you should know what happens when any of these values are changed. Beyond just changing a 1.3 to 2.3, you should also try changing 1.3 to 0.13, 130, and -1.3. Some of these values will produce bizarre and exciting results.

Sometimes it's even useful to automate the exploration of these values. Code up a matrix of values to try for each variable and let the program run all combinations. Then, curate the results and try out the promising leads.

When the program is nearly finished, you should experiment with slight changes to the values to make sure it's as polished as possible.

**Make the image stand on its own, without knowing the concept**

Visual artwork should communicate visually. If the viewer isn't interested by the image alone, you've failed here. The explanatory text should deepen the interest, but it is almost never more important than the image. If it is, you're really making conceptual art with a poor generative visualization. Maybe that's what you want, but for the purposes of this article, I'm assuming you want to make visual artwork.

If you have an interesting concept, do everything that you can to make the image reflect that level of awesomeness. You will catch the attention of the viewer with the image, and then you'll blow their mind with the concept. That's the order you want those things to happen in.

**Curate, don't post an album with 78 images**

If you write a program with randomized structure, it's going to produce a variety of output. This is one of the best parts of generative artwork! But don't make the mistake of thinking that all output is equal, or that anybody is interested in looking though more than ten of these. I think about this part of the work like a photographer. During a photo shoot, the photographer will collect hundreds of different images with lots of variety. The next step is to search through these to find the handful of images that really shine. Your job, as a generative artist, is to find the few images that highlight the strengths and versatility of your program.

With my own work, the number of images I choose to show depends on the program. When the program has low-variability output, I only show one image. If there's more variety in the output, I'll show about three. Use your own judgement here, but don't place the burden of doing this filtering on your viewer. Value the opportunity to select what images resonate for you.

**Print your artwork, frame it, and hang it**

This is most relevant if you're looking to sell your work. In my experience, selling generative artwork is easiest with this kind of presentation. I think it's hard for someone to make the leap from seeing a digital image on their phone screen to imagining it printed, framed, and on the wall in their home. But when you show somebody the already printed and framed work of art, it's very easy for them to imagine what it would look like on their wall. Generative artwork is already foreign to most viewers, so that imaginary gap needs to be as small as possible. Roughly 80% of my sales are to people that have seen my work in a physical format.

Even if you're not trying to sell anything, going through this process will allow you to see your work in a different context, which may help you to view it more objectively. This puts it in the realm of competing with other forms of visual artwork, and competition is good for encouraging you to improve.

One quick bit of advice: do not get cheap prints. Look for shops that print fine art. It's so worth the extra $20 it will cost you.

I hope these suggestions help you to improve your algorithmic artwork!

Cheers!