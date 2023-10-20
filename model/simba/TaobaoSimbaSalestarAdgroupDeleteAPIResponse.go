package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbasalestaradgroupdeleteAPIResponse (新)销量明星删除推广单元接口 API返回值
// taobao.simba.salestar.adgroup.delete
//
// 删除一个推广组
type TaobaosimbasalestaradgroupdeleteAPIResponse struct {
	model.CommonResponse
	TaobaosimbasalestaradgroupdeleteAPIResponseModel
}

// TaobaosimbasalestaradgroupdeleteAPIResponseModel is (新)销量明星删除推广单元接口 成功返回结果
type TaobaosimbasalestaradgroupdeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_salestar_adgroup_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 删除失败时透出的原因，仅当success为false是有效
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 表示接口调用成功，当推广单元不存在等情况下也会返回true
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
