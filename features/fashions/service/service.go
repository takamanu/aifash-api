package service

import "aifash-api/features/fashions"

type FashionService struct {
	fd fashions.FashionDataInterface
}

func NewService(fd fashions.FashionDataInterface) fashions.FashionServiceInterface {
	return &FashionService{
		fd: fd,
	}
}

func (fs *FashionService) StoreFashion(newData fashions.Fashion) (*fashions.Fashion, error) {

	res, err := fs.fd.StoreFashion(newData)

	if err != nil {
		return nil, err
	}

	return res, nil
}
func (fs *FashionService) GetAllFashion() ([]fashions.Fashion, error) {
	res, err := fs.fd.GetAllFashion()

	if err != nil {
		return nil, err
	}

	return res, nil
}
func (fs *FashionService) GetFashionByID(id int) (*fashions.Fashion, error) {
	res, err := fs.fd.GetFashionByID(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}
func (fs *FashionService) GetFashionByUserID(userID int) ([]fashions.Fashion, error) {
	res, err := fs.fd.GetFashionByUserID(userID)

	if err != nil {
		return nil, err
	}

	return res, nil
}
func (fs *FashionService) UpdateFashionByID(id int, newData fashions.Fashion) (bool, error) {
	_, err := fs.fd.UpdateFashionByID(id, newData)

	if err != nil {
		return false, err
	}

	return true, nil
}
func (fs *FashionService) DeleteFashionByID(id int) (bool, error) {
	_, err := fs.fd.DeleteFashionByID(id)

	if err != nil {
		return false, err
	}

	return true, nil
}
