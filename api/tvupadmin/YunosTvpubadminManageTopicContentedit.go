package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadminmanagetopiccontentedit 专题关联内容编辑
// yunos.tvpubadmin.manage.topic.contentedit
//
// 编辑专题关联的内容
func Yunostvpubadminmanagetopiccontentedit(clt *core.SDKClient, req *tvupadmin.YunostvpubadminmanagetopiccontenteditAPIRequest, session string) (*tvupadmin.YunostvpubadminmanagetopiccontenteditAPIResponse, error) {
	var resp tvupadmin.YunostvpubadminmanagetopiccontenteditAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
