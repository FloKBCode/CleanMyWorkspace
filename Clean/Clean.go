package Clean

func CleanWorkSpace(workspace *[][]*string) *[][]*string {
	if workspace == nil {
		return workspace
	}

	for i := range *workspace {
		for j := range (*workspace)[i] {
			(*workspace)[i][j] = nil
		}
	}

	return workspace
}
