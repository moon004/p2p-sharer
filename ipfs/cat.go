package ipfs

import (
	"context"
	"io"

	"github.com/ipfs/go-ipfs/core/coreapi/interface"

	"gx/ipfs/QmXWZCd8jfaHmt4UDSnjKmGcrQMw95bDGWqEeVLVJjoANX/go-ipfs-files"
)

// Cat will output hash content
func Cat(
	ctx context.Context,
	api iface.CoreAPI,
	path string,
	offset int64,
	max int64,
) (io.Reader, uint64, error) {
	var readers io.Reader
	length := uint64(0)
	if max == 0 {
		return nil, 0, nil
	}
	fpath, err := iface.ParsePath(path)
	if err != nil {
		return nil, 0, err
	}

	f, err := api.Unixfs().Get(ctx, fpath)
	if err != nil {
		return nil, 0, err
	}

	file, ok := f.(files.File)
	if !ok {
		return nil, 0, iface.ErrNotFile
	}

	fsize, err := file.Size()
	if err != nil {
		return nil, 0, err
	}

	if offset > fsize {
		offset = offset - fsize
	}

	count, err := file.Seek(offset, io.SeekStart)
	if err != nil {
		return nil, 0, err
	}
	offset = 0

	fsize, err = file.Size()
	if err != nil {
		return nil, 0, err
	}

	size := uint64(fsize - count)
	length += size
	if max > 0 && length >= uint64(max) {
		var r io.Reader = file
		if overshoot := int64(length - uint64(max)); overshoot != 0 {
			r = io.LimitReader(file, int64(size)-overshoot)
			length = uint64(max)
		}
		readers = r
	}
	readers = file

	return readers, length, nil
}
