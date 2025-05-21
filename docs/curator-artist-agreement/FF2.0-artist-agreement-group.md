# Feral File Artist Agreement - Group Exhibition

*Last Updated: 20-MAY 2025*

**Artist Name:** {{ .Signer.FullName }}

**Planned Launch Date:** {{ .OpeningAt.Format "02 January 2006 - 15:04 UTC" }}

You are invited to a group exhibition to be opened on Feral File, tentatively scheduled for the date above.

Below Feral File Inc. (“we” “us” “our” “Feral File”) provides you with the terms of our agreement for your work to be exhibited and sold on Feral File.

## Summary of Your Responsibilities:

1. **Complete the Catalogue Raisonné** through the Feral File Admin interface.  
   - Share career highlights \- press mentions, awards, notable installations \- and biographical information.  
   - Update complete artwork minting history.  
   - Grant Feral File rights to display and promote your past artworks through its website and apps.

2. **Create your artworks and provide exhibition sets if requested.**  
   Work with the curator to define one or more series & artworks to be exhibited.

**Artwork Requirements**

- Each artwork must be exclusive to this exhibition at the time of launch.  
- Artworks may build upon or be conceptually related to existing work, but must not have been previously sold or publicly exhibited elsewhere.

**Exhibition Sets**

- We may require that each artist create an equal number of series, each with the same number of artworks.  
- We may package series into “exhibition sets” for sale consisting of one artwork from each series from each artist.  
- Exhibition sets may be available for the entire exhibition or for a limited time, after which the artworks may be sold separately.  
- Artworks not in an exhibition set are sold separately.  
    
3. **Install your artwork(s)** through the Feral File Admin interface.  
     
4. **Prepare for the exhibition** with your curator.  
     
5. **Promote the exhibition.**  
   - Collaborate on promotional content, including interviews, essays, or video features.  
   - Promote the exhibition across your networks: social media, press contacts, existing collectors of your work, potential collectors, and mail lists.

6. **Support the financial success of the exhibition and your artworks.**  
   - Withhold from selling your artworks until exhibition sales are complete (all artworks sold out or removed from the primary market on Feral File).  
   - For 24 months after the end of the exhibition, Feral File will have the exclusive right to list unsold artworks on Feral File and mutually agreeable third-party platforms, subject to compensation terms and price agreements.

Specific due dates for all tasks are shared within the Feral File Admin interface. Failure to complete tasks on time may delay or result in removal of your participation in the exhibition.

## Compensation and Pricing Terms:

### Pricing

Feral File will suggest pricing for your artworks based on our experience in operating exhibitions .

- You may accept the suggested pricing or discuss with Feral File to work toward mutually agreeable pricing.  
- If an agreement on pricing cannot be reached, Feral File may remove your artworks from the exhibition.

### Primary Sale

- Feral File recoups any applicable transaction fees (typically around 3%, depending on the payment method).  
- After these fees, the remaining proceeds are split — {{ percent .FirstSale.Artist }} to you and {{ percent .FirstSale.Platform }} to Feral File.

### Secondary Sale

- You receive {{ percent .ReSale.Artist }} of the purchase price minus transaction fees.

### Exhibition Set of Artworks for Artists, Curator, and Feral File

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

**Company Name (if applicable):** 

**Artist Name (printed):** {{ .Signer.FullName }}

**Artist Signature:** {{ if .ExhibitionLegalSignature }} {{ .ExhibitionLegalSignature.Signature }} {{ end }}

**Date:** {{ if .ExhibitionLegalSignature }} {{ .ExhibitionLegalSignature.CreatedAt.Format "02 January 2006 - 15:04 UTC" }} {{ end }}
