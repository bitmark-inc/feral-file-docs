# Curator Agreement

**{{ .Signer.FullName }}:**

Bitmark, Inc. (“we”, “us”, “our”, “Feral File”) is providing you with the terms of the agreement for your work curating a new exhibition on Feral File to be launched on {{ .OpeningAt.Format "02 January 2006 - 15:04 UTC" }} with the responsibilities detailed in [The Curator’s Guide](https://docs.google.com/document/d/1g6F7IiDB5D3jWExJ8-HHfJdtZB5TlB-s-Zkuz5bztTI/edit?usp=sharing). There may be other tasks discussed and agreed upon by both parties (Feral File, yourself) over email. Those tasks, once agreed, will be part of this agreement.

**For these responsibilities, we agree to pay you the following compensation:**
- **Fixed fee of $2,500, paid in USDC no later than two months before the opening date of the exhibition, or [MONTH/DAY/YEAR] to [WALLET ADDRESS].**
- **{{ share .FirstSale }} share of Feral File’s portion of Net Sales from the exhibition**, paid in the currency of the transaction (USDC is used for credit card transactions), after Feral File has recouped from the artists’ share of the artwork sales all advances paid to the artists.
- **You will receive an artwork from each series or editioned artwork of a group exhibition or from each series or editioned artwork of a solo exhibition that is larger than 10 artworks (or as otherwise agreed).** Future sales or other dispositions of such editions are subject to the [Art Sale Agreement](https://feralfile.com/docs/art-sale-agreement) with the following fee structure: You will pay the artist {{ percent .ReSale.Artist }} of the selling price, pay Feral File {{ percent .ReSale.Platform }} of the selling price, and retain {{ percent .ReSale.Seller }} of the selling price, minus transaction fees. 

You can use [this worksheet](https://docs.google.com/spreadsheets/d/1n5rQ2CEga6c29YXaEiQXp1fkPLjfYJIws0W0sEhsLFo/edit#gid=1326931850) to understand how the payment structure other than the fixed fee breaks down (you will need to make a copy to edit the figures). Transaction fees may differ based on the collector’s payment method (cryptocurrency, credit card) and the exhibition’s minting fees, but typically average 3%.

You must withhold from selling any and all of the artworks allocated to you until exhibition sales are complete (all artworks sold out or removed from the primary market on Feral File).

**Feral File Overview**
- The idea of Feral File is that we will succeed together. 
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
   The Curator and Artists jointly promote the exhibition and each other
    - Join online for online openings and social events, such as live discussions through Twitter Spaces.
    - Be available for press interviews and other content production used to promote the exhibition.
    - Promote the exhibition across your networks: social media, press contacts, potential collectors, mail lists. (The Feral File team will work closely with you to maximize your impact and produce marketing assets for distribution.)
    - Artists provide a short biography for the curator to edit, sharing career highlights (press mentions, awards, notable artwork appearances/installations) with Feral File with such information used to promote artists and the exhibition.

- Group:
    - Curators invite artists to collaborate together as a collective.
    - All artists create an equal number of series or editioned artworks for the exhibition, each of the same number.
        - Ex. Artists A, B, C each create a series of 50 unique artworks. Artists D, E create artworks with 50 editions. Each artist has 50 items in the exhibition.
    - [–GRAPH](https://feralfile.com/exhibitions/graph-eg6) is a strong example of a model Feral File group exhibition
- Solo: 
    - Curators invite the artist to collaborate together on the exhibition.
    - Solo exhibitions feature multiple artworks (such as 1/1’s, series) that fit together under a theme.
    - The artist(s), curator, and Feral File receive one artwork from each series/artwork that is larger than 10 artworks / editions (or as otherwise agreed).
    - [Unsupervised](https://feralfile.com/exhibitions/unsupervised-sla) is a strong example of a model Feral File solo exhibition.

**In general, artwork installation and all curator deliverables must be completed <u>at least six weeks in advance of the opening<u>. Specific due dates for all tasks are shared within the Feral File exhibition installation interface.**

**Promoting the Exhibition and your Exhibition Artists**
1. Join online for online openings and social events, such as live discussions through Twitter Spaces.
2. Be available for press interviews and other content production used to promote the exhibition.
3. Promote the exhibition across your networks: social media, press contacts, potential collectors, mail lists. (The Feral File team will work closely with you to maximize your impact and produce marketing assets for distribution.)
4. Provide a short biography for the curator to edit. Share career highlights (press mentions, awards, notable artwork appearances/installations) with Feral File. This information will be used to promote you and the exhibition.

If you are curating this exhibition as part of your employment, you represent to Feral File: (a) that you are receiving all artworks on behalf of your employer, and that you agree to the Art Sale Agreement on your employer’s behalf and have the authority to do so, or (b) that you have your employer’s permission to receive all artworks on your own behalf and that you personally agree to the Art Sale Agreement.

You will receive sale proceeds directly in your preferred web3 wallet(s), which you must designate using the Feral File exhibition interface.

If the exhibition does not proceed due to events beyond the reasonable control of Feral File, or if you personally are not able to or do not fulfill your obligations as curator (e.g., if you are no longer employed by your current employer, or do not timely complete your deliverables), this Agreement will be terminated, and you will return the fixed curation fee immediately upon Feral File’s request. Except for the return of the fee, termination under this paragraph will be without liability or obligation for you or Feral File. Any rights in the artworks granted to you will not apply.

You understand and agree that this agreement contains all of Feral File’s obligations to you, and together with the documents referenced in this agreement, supersedes all prior understandings regarding its subject matter. You waive any claims against Feral File and its management and personnel in excess of the obligations set forth in this agreement and hold Feral File and its management and personnel harmless against any claims which may arise from your activities under this agreement. 

If you understand and accept the responsibilities explained above and in The Curator’s Guide, please confirm acceptance of this project and the terms of this agreement by signing below.

Understood and Agreed:

<br>Curator Name: {{ .Signer.FullName }}
{{ if .Signer.Metadata.Company }}<br>Employer (if any): {{ .Signer.Metadata.Company }}{{ end }}
{{ if .ExhibitionLegalSignature }}
<br>Signature: {{ .ExhibitionLegalSignature.Signature }}
<br>Date: {{ .ExhibitionLegalSignature.CreatedAt.Format "02 January 2006 - 15:04 UTC" }}
{{ end }}
