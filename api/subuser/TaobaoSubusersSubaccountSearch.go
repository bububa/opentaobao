package subuser

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/subuser"
)

// TaobaoSubusersSubaccountSearch 根据子账号登录名后缀模糊搜索子账号列表
// taobao.subusers.subaccount.search
//
// 根据子账号冒号后缀搜索子账号列表，支持中文单字、英文单词（不支持英文单字母） 分词规则搜索，该搜索词必传。模糊搜索使用阿里云搜索引擎所以该接口增值收费，如果不需要模糊搜索仅需要分页获取子账号列表，请使用taobao.sellercenter.subusers.page接口
func TaobaoSubusersSubaccountSearch(clt *core.SDKClient, req *subuser.TaobaoSubusersSubaccountSearchAPIRequest, session string) (*subuser.TaobaoSubusersSubaccountSearchAPIResponse, error) {
	var resp subuser.TaobaoSubusersSubaccountSearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
