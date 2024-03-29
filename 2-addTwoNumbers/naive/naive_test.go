package addTwoNumbers_naive

import (
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	t.Parallel()

	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example 1",
			args: args{
				l1: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}},
				l2: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}},
			},
			want: &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8, Next: nil}}},
		},
		{
			name: "Example 2",
			args: args{
				l1: &ListNode{Val: 0, Next: nil},
				l2: &ListNode{Val: 0, Next: nil},
			},
			want: &ListNode{Val: 0, Next: nil},
		},
		{
			name: "Example 3",
			args: args{
				l1: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}}}}},
				l2: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}},
			},
			want: &ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: nil}}}}}}}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			for got := addTwoNumbers(tt.args.l1, tt.args.l2); got != nil && tt.want != nil; got, tt.want = got.Next, tt.want.Next {
				if got.Val != tt.want.Val {
					t.Errorf("addTwoNumbers() = %v, want %v", got.Val, tt.want.Val)
				}
			}
		})
	}
}
