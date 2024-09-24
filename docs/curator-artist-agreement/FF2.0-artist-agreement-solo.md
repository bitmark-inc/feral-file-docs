# Artist Agreement Solo

**{{ .Signer.FullName }}:**
Curator {{ .Curator.FullName }} has invited you to a solo exhibition to be opened on Feral File in {{ .OpeningAt.Format "02 January 2006 - 15:04 UTC" }}.

Below, Bitmark, Inc. (“we”, “us”, “our”, “Feral File”) provides you with the terms of our agreement for your work to be exhibited and sold on Feral File.

Below is a summary of your responsibilities:

1. **Create your Artworks.**
   1. Work with the curator to define the theme for your exhibition and the artworks to be exhibited.
   1. This work must not be available anywhere else, but it can be related to existing work.
   1. Create a collection of thematically related artworks for the exhibition, which may be single or editioned artworks.

1. **Install your artwork(s) through the Feral File Admin interface**
   1. Prepare for the exhibition with your curator. (The Feral File team will assist with installation issues.)

1. **Promote the exhibition**
   1. Join online for online openings and social events, such as live discussions through Twitter Spaces.
   1. Be available for press interviews and other content production used to promote the exhibition.
   1. Promote the exhibition across your networks: social media, press contacts, existing collectors of your work, potential collectors, mail lists. (The Feral File team will work closely with you to maximize your impact and produce marketing assets for distribution.)
   1. Provide a short biography for the curator to edit. Share career highlights (press mentions, awards, notable artwork appearances/installations) with Feral File. This information will be used to promote you and the exhibition.

1. **Support the financial success of the exhibition and your artworks**
   1. Withhold from selling your artworks until exhibition sales are complete (all artworks sold out or removed from the primary market on Feral File).
   1. For 24 months after the end of the exhibition, Feral File will have the exclusive right to list on Feral File and mutually agreeable third party platforms any of your artworks unsold after the completion of the exhibition, subject to (i) the compensation terms below, and (ii) setting the price of your artwork at the price agreed between you and Feral File below, unless you and Feral File agree otherwise.


In general, artwork installation must be completed at least six weeks in advance of the opening. Specific due dates for all tasks are shared within the Feral File Admin interface. Failure to complete tasks on time may delay the exhibition or force the cancellation of the exhibition.

For these responsibilities, we agree to the following compensation and pricing terms:

1. **Pricing.**
   1. Feral File will suggest pricing for your artworks based on our experience in operating digital artwork exhibitions. You may accept the suggested pricing or discuss pricing with Feral File to work toward mutually agreeable pricing. If you or Feral File determine that such agreement is not possible, either you or Feral File may terminate this agreement and cancel the exhibition.
1. **Primary Sale**
   1. You receive {{ percent .FirstSale.Artist }} of the purchase price, minus any applicable transaction fees\* paid to the payment processor.
1. **Secondary Sale**
   1. You receive {{ percent .ReSale.Artist }} of the purchase price, minus transaction fees.
1. **Artworks for You, Curator and Feral File**
   1. At the conclusion of the NFT minting process, Feral File will automatically transfer artworks by mutual agreement free of charge.
   1. Your future sales or other dispositions of such artworks are subject to the [Agreement of Original Transfer of Work of Digital Art](https://feralfile.com/docs/art-sale-agreement{{ .QueryParams }}) in effect at the time of the exhibition, under which you will pay Feral File {{ percent .ReSale.Platform }} of the selling price, and will retain {{ percent .FirstSale.Seller }} of the selling price, minus transaction fees. (The smart contract governing secondary sales is designed to pay any seller {{ percent .ReSale.Seller }} of the selling price, and to pay the artist {{ percent .ReSale.Artist }}. So when you are the seller, you would receive {{ percent (add .ReSale.Seller .ReSale.Artist) }} in two separate payments.)

_\* Transaction fees may differ based on the collector’s payment method (cryptocurrency, credit card) and the exhibition’s minting fees, but typically average 3%._

You will receive sale proceeds directly in your preferred web3 wallet, which you must designate using the Feral File Admin interface.

If the exhibition does not proceed due to events beyond the reasonable control of Feral File, this agreement will be deemed canceled, without liability or obligation for you or Feral File. Any rights in your artworks granted to Feral File or the curator will not apply, and Feral File will securely destroy any copies of the artworks in its possession or control.

You can review the full terms and conditions of your rights in the current version of the [Terms for Artists and Collectors](https://feralfile.com/docs/terms-of-artist-collector{{ .QueryParams }}), and the [Agreement of Original Transfer of Work of Digital Art](https://feralfile.com/docs/art-sale-agreement{{ .QueryParams }}).

You understand and agree that this agreement, together with the documents referenced in this agreement, supersedes all prior understandings regarding its subject matter.

Referenced Documents:
- [Terms for Artists and Collectors](https://feralfile.com/docs/terms-of-artist-collector{{ .QueryParams }})
- [Agreement of Original Transfer of Work of Digital Art](https://feralfile.com/docs/art-sale-agreement{{ .QueryParams }})

Understood and Agreed:
<br>(Click "Sign contract" button below) (If you have assigned your rights in your artwork to a company you control, you may confirm that you control that company and and its acceptance of this project and the terms of this agreement by filling out the Employer field and clicking "Sign contract")
<br>Artist Name: {{ .Signer.FullName }}
{{ if .Signer.Metadata.Company }}
<br>Employer: {{ .Signer.Metadata.Company }}{{ end }}
{{ if .ExhibitionLegalSignature }}
<br>Signature: {{ .ExhibitionLegalSignature.Signature }}
<br>Date: {{ .ExhibitionLegalSignature.CreatedAt.Format "02 January 2006 - 15:04 UTC" }}
{{ end }}
