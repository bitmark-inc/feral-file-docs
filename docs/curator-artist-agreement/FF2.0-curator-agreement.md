# Curator Agreement

**{{ .Signer.AlumniAccount.FullName }}:**

Below, Bitmark, Inc. (“we”, “us”, “our”, “Feral File”) provides you with the terms of our agreement for your work curating a new exhibition on Feral File to be launched on {{ .OpeningAt.Format "02 January 2006 - 15:04 UTC" }} with the responsibilities detailed in [The Curator’s Guide](https://docs.google.com/document/d/1g6F7IiDB5D3jWExJ8-HHfJdtZB5TlB-s-Zkuz5bztTI/edit). There may be other tasks discussed and agreed upon by both parties (Feral File, yourself) over email. Those tasks, once agreed, will be part of this agreement.

**For these responsibilities, we agree to pay you the following compensation:**

- Fixed fee of $2,500, paid in ETH to {{ .Signer.ID }} by 30 days after signing this agreement.
- {{ share .FirstSale }} share of Feral File’s portion of net sales (amounts received less transaction fees) from the exhibition, after Feral File has recouped from the artists’ share of the artwork sales all advances paid to the artists.
- You will receive an artwork from each series or editioned artwork of a group exhibition or from each series or editioned artwork of a solo exhibition that is larger than 10 artworks (or as otherwise agreed). Future sales or other dispositions of such editions are subject to the [Agreement of Original Transfer of Work of Digital Art](https://feralfile.com/docs/art-sale-agreement{{ .QueryParams }}) in effect at the time of the exhibition with the following fee structure: You will pay the artist {{ percent .ReSale.Artist }} of the selling price, pay Feral File {{ percent .ReSale.Platform }} of the selling price, and retain {{ percent .ReSale.Seller }} of the selling price, minus transaction fees.

In general, artwork installation and all curator deliverables must be completed at least six weeks in advance of the opening. Specific due dates for all tasks are shared within the Feral File Admin interface.

Transaction fees may differ based on the collector’s payment method (cryptocurrency, credit card) and the exhibition’s minting fees, but typically average 3%.

You must withhold from selling any and all of the artworks allocated to you until exhibition sales are complete (all artworks sold out or removed from the primary market on Feral File).

If you are curating this exhibition as part of your employment, you represent to Feral File: (a) that you are receiving all artworks on behalf of your employer, and that you agree to the Art Sale Agreement on your employer’s behalf and have the authority to do so, or (b) that you have your employer’s permission to receive all artworks on your own behalf and that you personally agree to the Art Sale Agreement.

You will receive sale proceeds directly in your preferred web3 wallet, which you must designate using the Feral File exhibition interface.

If the exhibition does not proceed due to events beyond the reasonable control of Feral File, or if you personally are not able to or do not fulfill your obligations as curator (e.g., if you are no longer employed by your current employer, or do not timely complete your deliverables), this Agreement will be terminated, and you will return the fixed curation fee immediately upon Feral File’s request. Except for the return of the fee, termination under this paragraph will be without liability or obligation for you or Feral File. Any rights in the artworks granted to you will not apply.

You waive any claims against Feral File and its management and personnel in excess of the obligations set forth in this agreement and hold Feral File and its management and personnel harmless against any claims which may arise from your activities under this agreement.

You understand and agree that this agreement contains all of Feral File’s obligations to you, and together with the documents referenced in this agreement, supersedes all prior understandings regarding its subject matter.

**Referenced Documents:**

- [The Curator’s Guide](https://docs.google.com/document/d/1g6F7IiDB5D3jWExJ8-HHfJdtZB5TlB-s-Zkuz5bztTI/edit?usp=sharing)
- [Agreement of Original Transfer of Work of Digital Art](https://feralfile.com/docs/art-sale-agreement{{ .QueryParams }})

If you understand and accept the responsibilities explained above, please confirm acceptance of this project and the terms of this agreement by clicking “Sign contract” button below.

Understood and Agreed:
<br>Curator Name: {{ .Signer.AlumniAccount.FullName }}
{{ if .Signer.AlumniAccount.Company }}<br>Employer (if any): {{ .Signer.AlumniAccount.Company }}{{ end }}
{{ if .ExhibitionLegalSignature }}
<br>Signature: {{ .ExhibitionLegalSignature.Signature }}
<br>Date: {{ .ExhibitionLegalSignature.CreatedAt.Format "02 January 2006 - 15:04 UTC" }}
{{ end }}
