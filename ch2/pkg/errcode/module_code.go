package errcode

var (
	ErrorGetTagListFail = NewError(20010001, "ErrorGetTagListFail")
	ErrorCreateTagFail  = NewError(20010002, "ErrorCreateTagFail")
	ErrorUpdateTagFail  = NewError(20010003, "ErrorUpdateTagFail")
	ErrorDeleteTagFail  = NewError(20010004, "ErrorDeleteTagFail")
	ErrorCountTagFail   = NewError(20010005, "ErrorCountTagFail")

	ErrorGetArticleFail    = NewError(20020001, "ErrorGetArticleFail")
	ErrorGetArticlesFail   = NewError(20020002, "ErrorGetArticlesFail")
	ErrorCreateArticleFail = NewError(20020003, "ErrorCreateArticleFail")
	ErrorUpdateArticleFail = NewError(20020004, "ErrorUpdateArticleFail")
	ErrorDeleteArticleFail = NewError(20020005, "ErrorDeleteArticleFail")

	ErrorUploadFileFail = NewError(20030001, "ErrorUploadFileFail")
)
