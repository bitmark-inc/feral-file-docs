# Feral File Artist Agreement - Solo Exhibition

*Last Updated: 21-MAY 2025*

**{{ .Signer.FullName }}:**

**Planned Launch Date:** {{ .OpeningAt.Format "02 January 2006 - 15:04 UTC" }}

You are invited to a solo exhibition to be opened on Feral File, tentatively scheduled for the date above.

Below Feral File Inc. (“we” “us” “our” “Feral File”) provides you with the terms of our agreement for your work to be exhibited and sold on Feral File.

## Summary of Your Responsibilities:

1. **Complete the Catalogue Raisonné** through the Feral File Admin interface.  
   - Share career highlights – press mentions, awards, notable installations – and biographical information.  
   - Update complete artwork minting history.  
   - Grant Feral File rights to display and promote your past artworks through its website and apps.

2. **Create your Artworks.**  
   Work with the curator to define the theme for your exhibition and the artworks to be exhibited.  
   - Each artwork must be exclusive to this exhibition at the time of launch.  
   - Artworks may build upon or be conceptually related to existing work, but must not have been previously sold or publicly exhibited elsewhere.


3. **Install your artwork(s)** through the Feral File Admin interface.  
     
4. **Prepare for the exhibition** with your curator.  
     
5. **Promote the exhibition.**  
   - Collaborate on promotional content, including interviews, essays, or video features.  
   - Promote the exhibition across your networks: social media, press contacts, existing collectors, potential collectors, and mailing lists.


6. **Support the financial success of the exhibition and your artworks.**  
   - Refrain from offering these artworks for sale elsewhere until the exhibition concludes, defined as either all artworks being sold out or removed from the Feral File primary market.  
   - For 24 months after the end of the exhibition, Feral File will have the exclusive right to list unsold artworks on its platform and on mutually agreed third-party platforms, subject to compensation terms and pricing agreements.

Specific due dates for all tasks are shared within the Feral File Admin interface. Failure to complete tasks on time may delay or cancel the exhibition.

## Compensation and Pricing Terms:

Feral File will suggest pricing for your artworks based on our experience in operating exhibitions.

- You may accept the suggested pricing or discuss alternatives with us.  
- If pricing terms cannot be mutually agreed upon after good faith discussion, either party may withdraw from the exhibition without penalty.

### Primary Sale

- Feral File recoups any applicable transaction fees (typically around 3%, depending on the payment method).  
- After these fees, the remaining proceeds are split — {{ percent .FirstSale.Artist }} to you and {{ percent .FirstSale.Platform }} to Feral File.

### Secondary Sale

- You receive {{ percent .ReSale.Artist }} of the secondary sale price, after deducting any applicable transaction fees.

### Artworks for You, Curator, and Feral File

At the conclusion of the NFT minting process, Feral File will automatically transfer artworks by mutual agreement free of charge.

- Future sales or dispositions of such artworks are subject to the [Art Transfer Agreement](https://feralfile.com/legal/art-transfer-agreement) in effect at the time of the exhibition, under which you will pay Feral File {{ percent .ReSale.Platform }} of the selling price and retain {{ percent (add .ReSale.Seller .ReSale.Artist) }} minus transaction fees.
- (The smart contract governing secondary sales is designed to pay any seller {{ percent .ReSale.Seller }} of the selling price and the artist {{ percent .ReSale.Artist }}. When you are the seller, you receive {{ percent (add .ReSale.Seller .ReSale.Artist) }} in two separate payments.)

You will receive sale proceeds directly in your preferred web3 wallet, which you must designate using the Feral File Admin interface.

If the exhibition does not proceed due to events beyond the reasonable control of Feral File, this agreement will be deemed canceled without liability or obligation for you or Feral File. Any rights in your artworks granted to Feral File or the curator will not apply, and Feral File will securely destroy any copies of the artworks in its possession or control.

You can review the full terms and conditions of your rights in the current version of the [Terms for Artists and Collectors](https://www.feralfile.com/legal/terms-artists-and-collectors) and the [Art Transfer Agreement](https://feralfile.com/legal/art-transfer-agreement).

## Understood and Agreed:

By signing below, you confirm your acceptance of the terms outlined in this agreement.

If your artwork rights are assigned to a company you control, you may indicate that here and confirm you are authorized to sign on its behalf.

---

**Company Name (if applicable):** {{ if .Signer.Metadata.Company }} {{ .Signer.Metadata.Company }} {{ end }}

**Artist Name (printed):** {{ .Signer.FullName }}

**Artist Signature:** {{ if .ExhibitionLegalSignature }} {{ .ExhibitionLegalSignature.Signature }} {{ end }}

**Date:** {{ if .ExhibitionLegalSignature }} {{ .ExhibitionLegalSignature.CreatedAt.Format "02 January 2006 - 15:04 UTC" }} {{ end }}
