package downloader

import (
	. "base"
)

type IdGEenerTor interface {
	GetUInt32() uint32
}

type PageDownloader interface {
	Id() uint32
	Downloader(req base.Request) (*base.Response, error)
}
type PageDownloaderPool interface {
	Take() (PageDownloader, error)
	Return(dl PageDownloader) error
	Total() uint32
	Used() uint32
}
