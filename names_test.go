package names

import (
	"testing"
)

const targetTestVersion = 4

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v.", testVersion, targetTestVersion)
	}
}

func TestSeparateName(t *testing.T) {
	if err := LoadTitleData(); err != nil {
		t.Fatalf("Failed to load the prefix/suffix data and got the following error\n%s", err)
	}
	for _, ntc := range nameTestCases {
		got := New(ntc.input)
		got.SeparateName("")
		if got.First != ntc.expected.First {
			t.Fatalf("First Name ::: SeparateName(%q) = %q, want %q", ntc.input, got.First, ntc.expected.First)
		}
		if got.Middle != ntc.expected.Middle {
			t.Fatalf("Middle Name ::: SeparateName(%q) = %q, want %q", ntc.input, got.Middle, ntc.expected.Middle)
		}
		if got.Last != ntc.expected.Last {
			t.Fatalf("Last Name ::: SeparateName(%q) = %q, want %q", ntc.input, got.Last, ntc.expected.Last)
		}
		if got.Prefix != ntc.expected.Prefix {
			t.Fatalf("Prefix ::: SeparateName(%q) = %q, want %q", ntc.input, got.Prefix, ntc.expected.Prefix)
		}
		if got.Suffix != ntc.expected.Suffix {
			t.Fatalf("Suffix ::: SeparateName(%q) = %q, want %q", ntc.input, got.Suffix, ntc.expected.Suffix)
		}
	}
}

func TestLoadTitleDataCSV(t *testing.T) {
	for _, tt := range csvTests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LoadTitleDataCSV(tt.arg); (err != nil) != tt.wantErr {
				t.Errorf("LoadTitleDataCSV() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInitials(t *testing.T) {
	type fields struct {
		First     string
		Middle    string
		Last      string
		Prefix    string
		Suffix    string
		full      string
		formatted string
		initials  string
	}
	type args struct {
		dots bool
	}
	initialsTests := []struct {
		name   string
		fields fields
		arg    bool
		want   string
	}{
		{
			name:   "John Moore",
			fields: fields{First: "John", Middle: "K.", Last: "Moore"},
			arg:    true,
			want:   "J.K.M.",
		},
		{
			name:   "John Paul Jones",
			fields: fields{First: "John", Middle: "Paul", Last: "Jones"},
			arg:    false,
			want:   "JPJ",
		},
	}
	for _, tt := range initialsTests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Name{
				First:     tt.fields.First,
				Middle:    tt.fields.Middle,
				Last:      tt.fields.Last,
				Prefix:    tt.fields.Prefix,
				Suffix:    tt.fields.Suffix,
				full:      tt.fields.full,
				formatted: tt.fields.formatted,
				initials:  tt.fields.initials,
			}
			n.FormatName()
			if got := n.Initials(tt.arg); got != tt.want {
				t.Errorf("Name.Initials() = %v, want %v", got, tt.want)
			}
			if alreadyCreated := n.Initials(tt.arg); alreadyCreated == "" {
				t.Error("Initials should already be assigned.")
			}
		})
	}
}

func BenchmarkSeparateName(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		New(nameTestCases[2].input).SeparateName(" ")
	}
}

func BenchmarkLoadTitleData(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LoadTitleData()
	}
}
