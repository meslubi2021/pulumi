// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const MyEnum = {
    Pi: 3.1415,
    Small: 1e-07,
} as const;

export type MyEnum = (typeof MyEnum)[keyof typeof MyEnum];