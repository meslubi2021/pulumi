// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export { FooArgs } from "./foo";
export type Foo = import("./foo").Foo;
export const Foo: typeof import("./foo").Foo = null as any;
utilities.lazyLoad(exports, ["Foo"], () => require("./foo"));

export { FuncWithAllOptionalInputsArgs, FuncWithAllOptionalInputsResult, FuncWithAllOptionalInputsOutputArgs } from "./funcWithAllOptionalInputs";
export const funcWithAllOptionalInputs: typeof import("./funcWithAllOptionalInputs").funcWithAllOptionalInputs = null as any;
export const funcWithAllOptionalInputsOutput: typeof import("./funcWithAllOptionalInputs").funcWithAllOptionalInputsOutput = null as any;
utilities.lazyLoad(exports, ["funcWithAllOptionalInputs","funcWithAllOptionalInputsOutput"], () => require("./funcWithAllOptionalInputs"));

export { ModuleTestArgs } from "./moduleTest";
export type ModuleTest = import("./moduleTest").ModuleTest;
export const ModuleTest: typeof import("./moduleTest").ModuleTest = null as any;
utilities.lazyLoad(exports, ["ModuleTest"], () => require("./moduleTest"));

export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));


// Export sub-modules:
import * as types from "./types";

export {
    types,
};

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "example:index:Foo":
                return new Foo(name, <any>undefined, { urn })
            case "example:index:moduleTest":
                return new ModuleTest(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("example", "index", _module)
pulumi.runtime.registerResourcePackage("example", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:example") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
