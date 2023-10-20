package miniappcloud

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappcloudstorelistfileAPIResponse 云存储根据文件名反查地址 API返回值
// taobao.miniapp.cloud.store.listfile
//
// 云存储中，根据文件名反查地址
type TaobaominiappcloudstorelistfileAPIResponse struct {
	model.CommonResponse
	TaobaominiappcloudstorelistfileAPIResponseModel
}

// TaobaominiappcloudstorelistfileAPIResponseModel is 云存储根据文件名反查地址 成功返回结果
type TaobaominiappcloudstorelistfileAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_cloud_store_listfile_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Files []File `json:"files,omitempty" xml:"files>file,omitempty"`
}
