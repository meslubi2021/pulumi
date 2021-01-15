// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test/testdata/simple-enum-schema/go/plant"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type RubberTree struct {
	pulumi.CustomResourceState

	Container plant.ContainerPtrOutput `pulumi:"container"`
	Farm      pulumi.StringPtrOutput   `pulumi:"farm"`
	Type      pulumi.StringOutput      `pulumi:"type"`
}

// NewRubberTree registers a new resource with the given unique name, arguments, and options.
func NewRubberTree(ctx *pulumi.Context,
	name string, args *RubberTreeArgs, opts ...pulumi.ResourceOption) (*RubberTree, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Farm == nil {
		args.Farm = pulumi.StringPtr("(unknown)")
	}
	var resource RubberTree
	err := ctx.RegisterResource("plant:tree/v1:RubberTree", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRubberTree gets an existing RubberTree resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRubberTree(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RubberTreeState, opts ...pulumi.ResourceOption) (*RubberTree, error) {
	var resource RubberTree
	err := ctx.ReadResource("plant:tree/v1:RubberTree", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RubberTree resources.
type rubberTreeState struct {
	Container *plant.Container `pulumi:"container"`
	Farm      *string          `pulumi:"farm"`
	Type      *string          `pulumi:"type"`
}

type RubberTreeState struct {
	Container plant.ContainerPtrInput
	Farm      pulumi.StringPtrInput
	Type      *RubberTreeVariety
}

func (RubberTreeState) ElementType() reflect.Type {
	return reflect.TypeOf((*rubberTreeState)(nil)).Elem()
}

type rubberTreeArgs struct {
	Container *plant.Container `pulumi:"container"`
	Farm      *string          `pulumi:"farm"`
	Type      string           `pulumi:"type"`
}

// The set of arguments for constructing a RubberTree resource.
type RubberTreeArgs struct {
	Container plant.ContainerPtrInput
	Farm      pulumi.StringPtrInput
	Type      RubberTreeVariety
}

func (RubberTreeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rubberTreeArgs)(nil)).Elem()
}

type RubberTreeInput interface {
	pulumi.Input

	ToRubberTreeOutput() RubberTreeOutput
	ToRubberTreeOutputWithContext(ctx context.Context) RubberTreeOutput
}

func (*RubberTree) ElementType() reflect.Type {
	return reflect.TypeOf((*RubberTree)(nil))
}

func (i *RubberTree) ToRubberTreeOutput() RubberTreeOutput {
	return i.ToRubberTreeOutputWithContext(context.Background())
}

func (i *RubberTree) ToRubberTreeOutputWithContext(ctx context.Context) RubberTreeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RubberTreeOutput)
}

func (i *RubberTree) ToRubberTreePtrOutput() RubberTreePtrOutput {
	return i.ToRubberTreePtrOutputWithContext(context.Background())
}

func (i *RubberTree) ToRubberTreePtrOutputWithContext(ctx context.Context) RubberTreePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RubberTreePtrOutput)
}

type RubberTreePtrInput interface {
	pulumi.Input

	ToRubberTreePtrOutput() RubberTreePtrOutput
	ToRubberTreePtrOutputWithContext(ctx context.Context) RubberTreePtrOutput
}

type rubberTreePtrType RubberTreeArgs

func (*rubberTreePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RubberTree)(nil))
}

func (i *rubberTreePtrType) ToRubberTreePtrOutput() RubberTreePtrOutput {
	return i.ToRubberTreePtrOutputWithContext(context.Background())
}

func (i *rubberTreePtrType) ToRubberTreePtrOutputWithContext(ctx context.Context) RubberTreePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RubberTreePtrOutput)
}

// RubberTreeArrayInput is an input type that accepts RubberTreeArray and RubberTreeArrayOutput values.
// You can construct a concrete instance of `RubberTreeArrayInput` via:
//
//          RubberTreeArray{ RubberTreeArgs{...} }
type RubberTreeArrayInput interface {
	pulumi.Input

	ToRubberTreeArrayOutput() RubberTreeArrayOutput
	ToRubberTreeArrayOutputWithContext(context.Context) RubberTreeArrayOutput
}

type RubberTreeArray []RubberTreeInput

func (RubberTreeArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*RubberTree)(nil))
}

func (i RubberTreeArray) ToRubberTreeArrayOutput() RubberTreeArrayOutput {
	return i.ToRubberTreeArrayOutputWithContext(context.Background())
}

func (i RubberTreeArray) ToRubberTreeArrayOutputWithContext(ctx context.Context) RubberTreeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RubberTreeArrayOutput)
}

// RubberTreeMapInput is an input type that accepts RubberTreeMap and RubberTreeMapOutput values.
// You can construct a concrete instance of `RubberTreeMapInput` via:
//
//          RubberTreeMap{ "key": RubberTreeArgs{...} }
type RubberTreeMapInput interface {
	pulumi.Input

	ToRubberTreeMapOutput() RubberTreeMapOutput
	ToRubberTreeMapOutputWithContext(context.Context) RubberTreeMapOutput
}

type RubberTreeMap map[string]RubberTreeInput

func (RubberTreeMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*RubberTree)(nil))
}

func (i RubberTreeMap) ToRubberTreeMapOutput() RubberTreeMapOutput {
	return i.ToRubberTreeMapOutputWithContext(context.Background())
}

func (i RubberTreeMap) ToRubberTreeMapOutputWithContext(ctx context.Context) RubberTreeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RubberTreeMapOutput)
}

type RubberTreeOutput struct {
	*pulumi.OutputState
}

func (RubberTreeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RubberTree)(nil))
}

func (o RubberTreeOutput) ToRubberTreeOutput() RubberTreeOutput {
	return o
}

func (o RubberTreeOutput) ToRubberTreeOutputWithContext(ctx context.Context) RubberTreeOutput {
	return o
}

func (o RubberTreeOutput) ToRubberTreePtrOutput() RubberTreePtrOutput {
	return o.ToRubberTreePtrOutputWithContext(context.Background())
}

func (o RubberTreeOutput) ToRubberTreePtrOutputWithContext(ctx context.Context) RubberTreePtrOutput {
	return o.ApplyT(func(v RubberTree) *RubberTree {
		return &v
	}).(RubberTreePtrOutput)
}

type RubberTreePtrOutput struct {
	*pulumi.OutputState
}

func (RubberTreePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RubberTree)(nil))
}

func (o RubberTreePtrOutput) ToRubberTreePtrOutput() RubberTreePtrOutput {
	return o
}

func (o RubberTreePtrOutput) ToRubberTreePtrOutputWithContext(ctx context.Context) RubberTreePtrOutput {
	return o
}

type RubberTreeArrayOutput struct{ *pulumi.OutputState }

func (RubberTreeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RubberTree)(nil))
}

func (o RubberTreeArrayOutput) ToRubberTreeArrayOutput() RubberTreeArrayOutput {
	return o
}

func (o RubberTreeArrayOutput) ToRubberTreeArrayOutputWithContext(ctx context.Context) RubberTreeArrayOutput {
	return o
}

func (o RubberTreeArrayOutput) Index(i pulumi.IntInput) RubberTreeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RubberTree {
		return vs[0].([]RubberTree)[vs[1].(int)]
	}).(RubberTreeOutput)
}

type RubberTreeMapOutput struct{ *pulumi.OutputState }

func (RubberTreeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RubberTree)(nil))
}

func (o RubberTreeMapOutput) ToRubberTreeMapOutput() RubberTreeMapOutput {
	return o
}

func (o RubberTreeMapOutput) ToRubberTreeMapOutputWithContext(ctx context.Context) RubberTreeMapOutput {
	return o
}

func (o RubberTreeMapOutput) MapIndex(k pulumi.StringInput) RubberTreeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RubberTree {
		return vs[0].(map[string]RubberTree)[vs[1].(string)]
	}).(RubberTreeOutput)
}

func init() {
	pulumi.RegisterOutputType(RubberTreeOutput{})
	pulumi.RegisterOutputType(RubberTreePtrOutput{})
	pulumi.RegisterOutputType(RubberTreeArrayOutput{})
	pulumi.RegisterOutputType(RubberTreeMapOutput{})
}
