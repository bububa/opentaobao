package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
导入数据到商品百科服务 APIResponse
taobao.baike.import.zhubao.data

用于接入外部数据录入到商品百科中
*/
type TaobaoBaikeImportZhubaoDataAPIResponse struct {
    model.CommonResponse
    TaobaoBaikeImportZhubaoDataResponse
}

type TaobaoBaikeImportZhubaoDataResponse struct {
    XMLName xml.Name `xml:"baike_import_zhubao_data_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TaobaoBaikeImportZhubaoDataResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
