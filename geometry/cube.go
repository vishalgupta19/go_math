package geometry

func CubeVolume(n int) (int ,error) {

	if n!=0 {
		return n*n*n , nil
	}
	return 0, errors.new("Zero length is not allowed")
}