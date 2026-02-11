package greet

import "testing"

func TestMessage(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "with name",
			input: "Youta",
			want:  "Hello, Youta!",
		},
		{
			name:  "empty name uses default",
			input: "",
			want:  "Hello, Gopher!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Message(tt.input)
			if got != tt.want {
				t.Fatalf("Message(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
