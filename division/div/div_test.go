package div

import "testing"

func TestDivision(t *testing.T) {
	_, err := Division(8, 4)

	if err != nil {

		t.Error(err)
		return

	}

}

func TestDivisionErrorbyZero(t *testing.T) {
	_, err := Division(8, 0)

	if err == nil {

		t.Error(err)
		return

	}
}
