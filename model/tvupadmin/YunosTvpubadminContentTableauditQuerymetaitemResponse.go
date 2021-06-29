package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
运营位管控-查询魔盒运营位元数据列表 API返回值 
yunos.tvpubadmin.content.tableaudit.querymetaitem

运营位管控-查询魔盒运营位元数据列表
*/
type YunosTvpubadminContentTableauditQuerymetaitemAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminContentTableauditQuerymetaitemResponse
}

// 运营位管控-查询魔盒运营位元数据列表 成功返回结果
type YunosTvpubadminContentTableauditQuerymetaitemResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_content_tableaudit_querymetaitem_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 查看桌面坑位元数据列表（用于魔盒）
    Object   string `json:"object,omitempty" xml:"object,omitempty"`
}
