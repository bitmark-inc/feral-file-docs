# Artist Agreement

**{{ .Signer.FullName }}:**

Curator {{ .Curator.FullName }} has invited you to a group exhibition to be opened on Feral File in {{ .OpeningAt.Format "02 January 2006 - 15:04 UTC" }}. 

Below, Bitmark, Inc. (“we”, “us”, “our”, “Feral File”) provides  you with the terms of our agreement for your work to be exhibited and sold on [Feral File](https://feralfile.com). Please let us know if you have any questions.

**Feral File Overview**
- **The idea of Feral File is that we will succeed together. [–GRAPH](https://feralfile.com/exhibitions/graph-eg6) and [N=12](https://www.youtube.com/watch?v=jyTv3GLb4Ss&list=PLBtqCEMsHOxO4Yp0RvlRWbz_lKIAqOi-O&index=13) are strong examples of model Feral File group exhibitions.**
- Exhibitions start with a curator and their vision. Curators invite artists to collaborate together as a collective, driving the success of the exhibition, from its execution to its critical reception and financial outcomes.
- Each artist creates at least one series or editioned artwork for the exhibition. Each “item” in a series, a collection of related artworks, is a unique artwork.
- The exhibition will present a new series or editioned artwork from each artist.
- Artworks are “born-digital” and represent an ambitious creative effort by the artist. 
- All artists create an equal number of series or editioned artworks for the exhibition, each of the same number.
    - Ex. Artists A, B, C each create a series of 50 unique artworks. Artists D, E create artworks with 50 editions. Each artist has 50 items in the exhibition.
- Series can be
    - A few to thousands of artworks; and
    - created from any medium. For example, a series can be a single data type selected from: generative (code), video, image, 3d models, language, etc. 
- An “exhibition set” consists of one artwork (from series) or edition from each artist.
- Artists exhibit and collect the exhibition together. Each artist, the curator, and Feral File collects one exhibition set. 
- Artworks are sold to the public through a mix of exhibition sets and individual artworks/editions.

**In general, artwork installation must be completed <u>at least six weeks in advance of the opening<u>. Specific due dates for all tasks are shared within the Feral File exhibition installation interface.**

Failure to complete tasks on time may delay the exhibition or force the exhibition to continue without your participation.

Below is a summary of your responsibilities: 

1. **Install your artwork(s) through the Feral File exhibition installation interface**
    1.Work with the curator to define one or more series & artworks to be exhibited. 
    1. This work must not be available anywhere else, but it can be related to existing work.
    1. Prepare for the exhibition with your curator and other artists on Feral File’s Discord. (The Feral File team will assist with installation issues.)

1. **Promote the exhibition and your fellow exhibition artists**
    1. Join online for online openings and social events, such as live discussions through Twitter Spaces.
    1. Be available for press interviews and other content production used to promote the exhibition.
    1. Promote the exhibition across your networks: social media, press contacts, potential collectors, mail lists. (The Feral File team will work closely with you to maximize your impact and produce marketing assets for distribution.)
    1. Provide a short biography for the curator to edit. Share career highlights (press mentions, awards, notable artwork appearances/installations) with Feral File. This information will be used to promote you and the exhibition.

1. **Support the financial success of the exhibition and your artworks**
    1. Through the Feral File exhibition installation interface, transfer artworks (all artworks will be preminted in your name and then held for sale by Feral File) to the following persons free of charge, in the order listed:
        1. The curator of the exhibition;
        1. The other artists in the exhibition; and
        1. Feral File.
    1. Withhold from selling your exhibition set (in its entirety or through individual artworks) until exhibition sales are complete (all artworks sold out or removed from the primary market on Feral File).
    1. For 24 months after the completion of the exhibition, allow Feral File to list on mutually agreeable third party platforms any of your artworks still unsold after the completion of the exhibition, under the following conditions:
        1. the list pricing on the third party platforms will not be lower than the price agreed for the exhibition;
        1. the compensation terms below will continue to apply; and
        1. during this 24 month period you will have the right to reclaim the artworks by returning any of the advance paid under this Agreement which was not recouped by Feral File from sales of your artworks during the exhibition. 

If artworks are not automatically withdrawn from listings by the operation of a smart contract, you and Feral File will each inform the other of any sales of such artworks to allow withdrawal from other platforms.

**For these responsibilities, we agree to the following compensation and pricing terms:**
1. **Pricing**
    1. Feral File will propose pricing for your exhibition set with the goal of maximum sales for the exhibition as a whole, taking into account the estimated demand for your artworks, the content of the exhibition overall, and the correlation of pricing and sales results for past exhibitions.  You may accept the pricing proposal, and your exhibition set will be priced accordingly.  If you disagree with the proposed pricing, you may provide a counterproposal with supporting information, and you and Feral File will negotiate in good faith to reach agreed pricing.  If Feral File determines that such agreement is not possible, it may terminate this agreement, withdraw your artworks from the exhibition, and you will be free to list such artworks elsewhere. If you determine that such agreement is not possible, you may return the full advance paid under this Agreement, and upon repayment  withdraw your artworks from the exhibition, and you will be free to list such artworks elsewhere. In either case, notice must be provided by the terminating party (and for you, refund of the advance received) not later than 2 months prior to the exhibition opening date, or if this agreement is executed less than two months before the exhibition opening date, no later than 2 weeks after the Agreement execution date.
1. **Advance**
    1. Feral File will pay you an advance against sales of $2500, **paid in USDC to [WALLET ADDRESS] no later than two months before the opening date of the exhibition, or if this agreement is executed less than two months before the exhibition opening date, no later than [MMDDYYYY]**. You agree to return the full advance immediately upon Feral File’s request if the exhibition is canceled, if Feral File terminates this agreement based on a failure to reach agreement on pricing, or if you do not fulfill your obligations under this agreement, including without limitation your obligation to timely complete the installation of your exhibition set. 
1. **Primary Sale**
    1. You receive {{ percent .FirstSale.Artist }} of the purchase price, minus any applicable transaction fees* paid to the payment processor (the payment processor is not Feral File, please see note below). Feral File will credit the amounts due under this paragraph against the advance until the advance has been recouped, after which Feral File will pay you the amounts due after which Feral File will pay you the amounts due per the terms outlined in this clause. 
1. **Secondary Sale**
    1. You receive {{ percent .ReSale.Artist }} of the purchase price, minus transaction fees. If the full advance has not been recouped from revenue from the primary sale, Feral File will credit the amounts due under this paragraph against the advance until the advance has been recouped, after which Feral File will pay you the amounts due after which Feral File will pay you the amounts due per the terms outlined in this clause.
1. **Exhibition Set of Artworks from Artist Swaps**
    1. Each artist will produce artworks that are part of the “artist swap” for that exhibition. As noted above, you will receive an exhibition set of one artwork from each series or editioned artwork in the exhibition, including your own.
    2. Future sales or other dispositions of such artworks are subject to the [Art Sale Agreement](https://feralfile.com/docs/art-sale-agreement), under which you will pay the artist {{ percent .ReSale.Artist }} of the selling price, pay Feral File {{ percent .ReSale.Platform }} of the selling price, and will retain {{ percent .ReSale.Seller }} of the selling price, minus transaction fees. If Feral File receives payments of the full sale price from the secondary market, it will make distribution to you and the artist.  If you receive some or all of the sale price from the secondary market, you are responsible for ensuring  correct and full distribution to Feral File and the artist, including coordination with Feral File if necessary.

* Transaction fees may differ based on the collector’s payment method (cryptocurrency, credit card)  and the exhibition’s minting fees, but typically average 3%. You can use our worksheet to understand how the payment structure breaks down (you will need to make a copy to edit the figures): [Feral File Exhibition Financials - Artist + Curator - 2024](https://docs.google.com/spreadsheets/d/1n5rQ2CEga6c29YXaEiQXp1fkPLjfYJIws0W0sEhsLFo/edit#gid=1326931850) 

**You will receive sale proceeds directly in your preferred web3 wallet(s), which you must designate using the Feral File exhibition interface.**

If the exhibition does not proceed due to events beyond the reasonable control of Feral File, this Agreement will be deemed canceled, without liability or obligation for you or Feral File. Any rights in your artworks granted to Feral File, the curator, or other artists will not apply, and  Feral File  will securely destroy any copies of the artworks in its possession or control.

You can review the full terms and conditions of your rights in the current version of the [Art Sale Agreement](https://feralfile.com/docs/art-sale-agreement) and [Terms for Artists and Collectors](https://feralfile.com/docs/terms-of-artist-collector). You understand and agree that this agreement, together with the documents referenced in this agreement, supersedes all prior understandings regarding its subject matter.

Understood and Agreed:
<br>(Click “Sign contract” button below)
If you have assigned your rights in your artwork to a company you control, you may confirm that you control that company and its acceptance of this project and the terms of this agreement by filling out the Employer field and clicking “Sign contract”)

(You may also sign by clicking “Sign contract” button below)
{{ if .Signer.Metadata.Company }}(If you have assigned your rights in your artwork to a company you control, you may confirm that you control that company and its acceptance of this project and the terms of this agreement by filling out the Employer field and clicking “Sign contract”)
{{ end }}
<br>Artist Name: {{ .Signer.FullName }}
{{ if .Signer.Metadata.Company }}<br>Employer (if any): {{ .Signer.Metadata.Company }}{{ end }}
{{ if .ExhibitionLegalSignature }}
<br>Signature: {{ .ExhibitionLegalSignature.Signature }}
<br>Date: {{ .ExhibitionLegalSignature.CreatedAt.Format "02 January 2006 - 15:04 UTC" }}
{{ end }}
