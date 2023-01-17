package metadata

import (
	"strconv"

	"github.com/bububa/opentaobao/metadata/util"
)

const (
	// DOC_CENTER_URL 档中心链接，用于获取_tb_token_ cookie
	DOC_CENTER_URL = "https://open.taobao.com/docCenter"

	// API_CATELOG_URL 文档类目地址
	API_CATELOG_URL = "https://open.taobao.com/handler/document/getApiCatelogConfig.json?scopeId=&treeId=&docId=285&docType=2&tag=&_tb_token_="
)

// DocURL 单API文档地址
func DocURL(docID int64, tbToken string) string {
	return util.StringsJoin("https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=", strconv.FormatInt(docID, 10), "&docType=2&_tb_token_=", tbToken)
}

// DocLink API文档链接
func DocLink(docID int64) string {
	return util.StringsJoin("https://open.taobao.com/API.htm?docType=2&docId=", strconv.FormatInt(docID, 10))
}
