# Feral File Exhibition Curator Agreement

*Last Updated: 21-MAY 2025*

**Curator Name:** {{ .Signer.FullName }}

**Planned Launch Date:** {{ .OpeningAt.Format "02 January 2006 - 15:04 UTC" }}

Below Feral File Inc. (“we” “us” “our” “Feral File”) provides you with the terms of our agreement for your work curating a new exhibition on Feral File with the responsibilities detailed in [The Curator’s Guide](https://docs.google.com/document/d/12xfm1tAaFsJE7Loj3npa7Sag_QjWpIjjdZnBjY80DEs/edit?usp=sharing). There may be other tasks discussed and agreed upon by both parties (Feral File and yourself) over email. Those tasks, once agreed, will be part of this agreement.

## Compensation

For these responsibilities, we agree to pay you the following compensation:

- **Fixed fee** of $2500 paid in ETH within 30 days of the exhibition’s opening.  
- **{{ share .FirstSale }} share** of Feral File’s portion of net sales (amounts received less transaction fees) from the exhibition.  
- You will receive an **artwork(s)** as agreed with the artist and Feral File. Future sales or other dispositions of such editions are subject to the [Art Transfer Agreement](https://feralfile.com/legal/art-transfer-agreement) in effect at the time of the exhibition with the following fee structure:  
  - Pay the artist {{ percent .ReSale.Artist }} of the selling price
  - Pay Feral File {{ percent .ReSale.Platform }} of the selling price
  - Retain {{ percent .ReSale.Seller }} of the selling price minus transaction fees

### Additional Terms

- Artwork installation and all curator deliverables must be completed according to the dates set within the Feral File Admin interface. Failure to complete tasks on time may delay or cancel the exhibition.  
- Transaction fees may differ based on the collector’s payment method (cryptocurrency, credit card) and the exhibition’s minting fees but typically average 3%.  
- You must **withhold from selling any and all of the artworks allocated to you** until exhibition sales are complete (all artworks sold out or removed from the primary market on Feral File).

## Employment Representation

If you are curating this exhibition as part of your employment, you represent to Feral File that:

- (a) You are receiving all artworks on behalf of your employer, and you agree to the [Art Sale Agreement](https://feralfile.com/legal/art-sale-agreement) on your employer’s behalf, and have the authority to do so.  
- (b) You have your employer’s permission to receive all artworks on your own behalf, and you personally agree to the Art Sale Agreement.

## Payment and Termination

- You will receive sale proceeds directly in your preferred web3 wallet, which you must designate using the Feral File exhibition interface.  
- If the exhibition does not proceed due to events beyond the reasonable control of Feral File, or if you personally are unable to fulfill your obligations as curator (e.g., you are no longer employed by your current employer or do not complete your deliverables on time), this Agreement will be terminated, and you will return the fixed curation fee immediately upon Feral File’s request. Except for the return of the fee, termination under this paragraph will be without liability or obligation for you or Feral File. Any rights in the artworks granted to you will not apply.

## Liability Waiver

You waive any claims against Feral File and its management and personnel in excess of the obligations set forth in this agreement and hold Feral File and its management and personnel harmless against any claims which may arise from your activities under this agreement.

## Agreement and Signature

This document outlines the full agreement between you and Feral File regarding your role as curator. It supersedes any prior understandings and includes all referenced documents linked above.

By signing below, you confirm that you:

- Understand and accept the responsibilities outlined in this agreement and [The Curator’s Guide](https://docs.google.com/document/d/12xfm1tAaFsJE7Loj3npa7Sag_QjWpIjjdZnBjY80DEs/edit?usp=sharing).  
- Agree to the compensation terms, sales conditions, and artwork transfer terms.  
- (If applicable) Are authorized to sign on behalf of your employer or institution.

---


**Company Name (if applicable):** {{ if .Signer.Metadata.Company }} {{ .Signer.Metadata.Company }} {{ end }}

**Curator Name (printed):** {{ .Signer.FullName }}

**Curator Signature:** {{ if .ExhibitionLegalSignature }} {{ .ExhibitionLegalSignature.Signature }} {{ end }}

**Date:** {{ if .ExhibitionLegalSignature }} {{ .ExhibitionLegalSignature.CreatedAt.Format "02 January 2006 - 15:04 UTC" }} {{ end }}

