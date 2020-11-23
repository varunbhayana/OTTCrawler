package enum

type CrawlingIds string

func (id CrawlingIds) ToString() string {
	return string(id)

}

const (
	ZEE_5 CrawlingIds = "zee_5"
)
