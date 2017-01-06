package wallclock

import "testing"

func TestMinutes(t *testing.T) {
    cases := []struct {
        minute int
        want []string
    }{
        {5, []string{"five", "past"}},
    }
    for _, c := range cases {
        got, err := Minutes(c.minute)
        if err != nil {
            t.Error(err)
        }
        if !StringArrayEquals(c.want, got) {
            t.Errorf("Minutes(%d) == %q, want %q", c.minute, got, c.want)
        }
    }
}

func StringArrayEquals(a []string, b []string) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}
