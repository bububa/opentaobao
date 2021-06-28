package eticket

// QueryTbMaCallbackResp 
/* model for simplify = false
type QueryTbMaCallbackResp struct {

    // certificateDTO
    
    Certificate  *struct {
        CertificateDto  *CertificateDto `json:"certificate_dto,omitempty"`
    } `json:"certificate,omitempty"`
    

}
*/

// QueryTbMaCallbackResp 
type QueryTbMaCallbackResp struct {

    // certificateDTO
    Certificate   *CertificateDto `json:"certificate,omitempty"`

}
