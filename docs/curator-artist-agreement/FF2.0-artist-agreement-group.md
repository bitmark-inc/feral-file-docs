# Feral File Artist Agreement - Group Exhibition

**Last Updated: 09-JAN 2024**

**{{ .Signer.FullName }}:**

You are invited to a group exhibition to be opened on Feral File in {{ .OpeningAt.Format "02 January 2006 - 15:04 UTC" }}.

Below Bitmark Inc. (“we” “us” “our” “Feral File”) provides you with the terms of our agreement for your work to be exhibited and sold on Feral File.

## Summary of Your Responsibilities:

1. **Complete the Catalogue Raisonné** through the Feral File Admin interface.
Work with the curator to define the theme for your exhibition and the artworks to be exhibited.
    - Share career highlights - press mentions, awards, notable installations - and biographical information.
    - Update complete artwork minting history.
    - Grant Feral File rights to display and promote your past artworks through its website and apps.
  
1. **Create your Artworks.**
    - Work with the curator to define the theme for your exhibition and the artworks to be exhibited.
    - This work must not be available anywhere else but can be related to existing work.
    - Create a collection of thematically related artworks for the exhibition.

1. **Install your artwork(s)** through the Feral File Admin interface.
   
1. **Prepare for the exhibition** with your curator.

1. **Promote the exhibition.**
    - Be available for content production used to promote the exhibition.
    - Promote the exhibition across your networks: social media, press contacts, existing collectors of your work, potential collectors, and mail lists.

1. **Support the financial success of the exhibition and your artworks.**
    - Withhold from selling your artworks until exhibition sales are complete (all artworks sold out or removed from the primary market on Feral File).
    - For 24 months after the end of the exhibition, Feral File will have the exclusive right to list unsold artworks on Feral File and mutually agreeable third-party platforms, subject to compensation terms and price agreements.

Specific due dates for all tasks are shared within the Feral File Admin interface. Failure to complete tasks on time may delay or cancel the exhibition.

## Compensation and Pricing Terms:

### Pricing

Feral File will suggest pricing for your artworks based on our experience in operating exhibitions .
- You may accept the suggested pricing or discuss with Feral File to work toward mutually agreeable pricing.
- If an agreement cannot be reached, either you or Feral File may terminate this agreement and cancel the exhibition.

### Primary Sale
* Feral File recoups any applicable transaction fees.
* After recouping fees, you receive {{ percent .FirstSale.Artist }} of the purchase price paid to the payment processor.
* Feral File will credit the amounts due under this paragraph against the advance until recouped, after which you will be paid per the terms.

### Secondary Sale
* You receive {{ percent .ReSale.Artist }} of the purchase price minus transaction fees.

### Artworks for You, Curator, and Feral File
At the conclusion of the NFT minting process, Feral File will automatically transfer artworks by mutual agreement free of charge.

* Future sales or dispositions of such artworks are subject to the [Art Transfer Agreement](https://feralfile.com/legal/art-transfer-agreement) in effect at the time of the exhibition, under which you will pay Feral File {{ percent .ReSale.Platform }} of the selling price and retain {{ percent (add .ReSale.Seller .ReSale.Artist) }} minus transaction fees.
* (The smart contract governing secondary sales is designed to pay any seller {{ percent .ReSale.Seller }} of the selling price and the artist {{ percent .ReSale.Artist }}. When you are the seller, you receive {{ percent (add .ReSale.Seller .ReSale.Artist) }} in two separate payments.)

_* Transaction fees may differ based on the collector’s payment method (cryptocurrency, credit card) and minting fees but typically average 3%._

You will receive sale proceeds directly in your preferred web3 wallet, which you must designate using the Feral File Admin interface.

If the exhibition does not proceed due to events beyond the reasonable control of Feral File, this agreement will be deemed canceled without liability or obligation for you or Feral File. Any rights in your artworks granted to Feral File or the curator will not apply, and Feral File will securely destroy any copies of the artworks in its possession or control.

You can review the full terms and conditions of your rights in the current version of the [Terms for Artists and Collectors](https://www.feralfile.com/legal/terms-artists-and-collectors), and the [Art Transfer Agreement](https://feralfile.com/legal/art-transfer-agreement).

## Understood and Agreed:
(Click "Sign contract" button below) (If you have assigned your rights in your artwork to a company you control, you may confirm that you control that company and its acceptance of this project and the terms of this agreement by filling out the Employer field and clicking "Sign contract")

<br>Artist Name: {{ .Signer.FullName }}
<br>Signature:{{ if .ExhibitionLegalSignature }} {{ .ExhibitionLegalSignature.Signature }} {{ end }}
<br>Date:{{ if .ExhibitionLegalSignature }} {{ .ExhibitionLegalSignature.CreatedAt.Format "02 January 2006 - 15:04 UTC" }} {{ end }}
