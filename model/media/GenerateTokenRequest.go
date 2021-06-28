package media

// GenerateTokenRequest 
/* model for simplify = false
type GenerateTokenRequest struct {

    // 请求策略
    
    UploadPolicy  *struct {
        UploadPolicy  *UploadPolicy `json:"upload_policy,omitempty"`
    } `json:"upload_policy,omitempty"`
    

}
*/

// GenerateTokenRequest 
type GenerateTokenRequest struct {

    // 请求策略
    UploadPolicy   *UploadPolicy `json:"upload_policy,omitempty"`

}
