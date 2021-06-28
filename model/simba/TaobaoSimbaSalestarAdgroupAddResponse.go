package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
(新)创建一个推广组 APIResponse
taobao.simba.salestar.adgroup.add

创建一个推广组
*/
type TaobaoSimbaSalestarAdgroupAddAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_salestar_adgroup_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 新增加的推广组
    
    Adgroup   *ADGroup `json:"adgroup,omitempty" xml:"