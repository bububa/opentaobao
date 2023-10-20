package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmincontentshowedit 媒资节目信息修改
// yunos.tvpubadmin.content.show.edit
//
// 供迎客松修改媒资节目信息
func Yunostvpubadmincontentshowedit(clt *core.SDKClient, req *tvupadmin.YunostvpubadmincontentshoweditAPIRequest, session string) (*tvupadmin.YunostvpubadmincontentshoweditAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmincontentshoweditAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
