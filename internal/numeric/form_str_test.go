package numeric

import "testing"

func TestFormStr(t *testing.T) {
	cases := []struct {
		numeric  Response
		args     []interface{}
		expected []string
	}{
		{
			numeric:  RplWelcome,
			args:     []interface{}{"ShadowNet", "Xena"},
			expected: []string{"Welcome to the ShadowNet Internet Relay Chat Network Xena"},
		},
		{
			numeric:  RplMyinfo,
			args:     []interface{}{"ariablaze.ponychat.net", "elemental-ircd-6.6.2+ponychat", "BCDGQRSVWZagilopswxz", "CDEFGIJKLMOPQSTabcdefghijklmnopqrstuvyz", "yabefhjklovqI"},
			expected: []string{"ariablaze.ponychat.net", "elemental-ircd-6.6.2+ponychat", "BCDGQRSVWZagilopswxz", "CDEFGIJKLMOPQSTabcdefghijklmnopqrstuvyz", "yabefhjklovqI"},
		},
		{
			numeric:  RplSavenick,
			args:     []interface{}{"Xena"},
			expected: []string{"Xena", "Nick collision, forcing nick change to your unique ID"},
		},
	}

	for _, cs := range cases {
		t.Run(cs.numeric.String(), func(t *testing.T) {
			res := cs.numeric.FormStr(cs.args...)
			t.Logf("res: %#v", res)
			if len(res) != len(cs.expected) {
				t.Fatalf("length of output different than expected")
			}

			for i := 0; i < len(res); i++ {
				if res[i] != cs.expected[i] {
					t.Fatalf("res[%d] != cs.expected[%d]: %s != %s", i, i, res[i], cs.expected[i])
				}
			}
		})
	}
}
