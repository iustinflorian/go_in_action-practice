package processor

type ID int64

type Batch struct {
	BatchSize int
	BatchElem []ID
}

func getBatch(ids []ID, batchSize int) Batch {
	size := min(len(ids), batchSize)
	newBatch := ids[0:size:size]

	return Batch{
		BatchSize: batchSize,
		BatchElem: newBatch,
	}
}

func DisplayBatches(ids []ID, batchSize int) []Batch {
	var newBatchList []Batch

	dst := make([]ID, len(ids))
	copy(dst, ids)

	for len(dst) > 0 {
		newBatch := getBatch(dst, batchSize)
		newBatchList = append(newBatchList, newBatch)
		dst = dst[newBatch.BatchSize:]
	}

	return newBatchList
}
