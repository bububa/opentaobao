package security

// StsUploadToken 
type StsUploadToken struct {

    // accessKeyId
    AccessKeyId   string `json:"access_key_id,omitempty"`

    // accessKeySecret
    AccessKeySecret   string `json:"access_key_secret,omitempty"`

    // bucketName
    BucketName   string `json:"bucket_name,omitempty"`

    // endPoint
    EndPoint   string `json:"end_point,omitempty"`

    // expiration
    Expiration   string `json:"expiration,omitempty"`

    // path
    Path   string `json:"path,omitempty"`

    // token
    Token   string `json:"token,omitempty"`

}
