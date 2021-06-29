package miniappcloud

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
云存储根据文件名反查地址 APIResponse
taobao.miniapp.cloud.store.listfile

云存储中，根据文件名反查地址
*/
type TaobaoMiniappCloudStoreListfileAPIResponse struct {
    model.CommonResponse
    TaobaoMiniappCloudStoreListfileResponse
}

type TaobaoMiniappCloudStoreListfileResponse struct {
    XMLName xml.Name `xml:"miniapp_cloud_store_listfile_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Files   []File `json:"files,omitempty" xml:"files>file,omitempty"`
    
    
}
