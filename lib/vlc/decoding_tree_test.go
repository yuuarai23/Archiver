package vlc

import (
	"reflect"
	"testing"
)

func Test_encodingTable_DecodingTree(t *testing.T) {
	tests := []struct {
		name string
		et   encodingTable
		want DecodingTree
	}{
		{
			name: "test 1",
			et: encodingTable{
				'a': "011",
				'b': "0000010",
				'z': "000000000000",
			},
			want: DecodingTree{
				Zero: &DecodingTree{
					One: &DecodingTree{
						One: &DecodingTree{
							Value: "a",
						},
					},
					Zero: &DecodingTree{
						Zero: &DecodingTree{
							Zero: &DecodingTree{
								Zero: &DecodingTree{
									One: &DecodingTree{
										Zero: &DecodingTree{
											Value: "b",
										},
									},
									Zero: &DecodingTree{
										Zero: &DecodingTree{
											Zero: &DecodingTree{
												Zero: &DecodingTree{
													Zero: &DecodingTree{
														Zero: &DecodingTree{
															Zero: &DecodingTree{
																Value: "z",
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "test 2",
			et: encodingTable{
				'a': "011",
				'e': "101",
				'h': "0011",
				's': "0101",
				't': "1001",
			},
			want: DecodingTree{
				Zero: &DecodingTree{
					One: &DecodingTree{
						One: &DecodingTree{
							Value: "a",
						},
						Zero: &DecodingTree{
							One: &DecodingTree{
								Value: "s",
							},
						},
					},
					Zero: &DecodingTree{
						One: &DecodingTree{
							One: &DecodingTree{
								Value: "h",
							},
						},
					},
				},
				One: &DecodingTree{
					Zero: &DecodingTree{
						One: &DecodingTree{
							Value: "e",
						},
						Zero: &DecodingTree{
							One: &DecodingTree{
								Value: "t",
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.et.DecodingTree(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodingTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
