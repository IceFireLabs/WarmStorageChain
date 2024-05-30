/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "blockcraftsman.warmstoragechain.warmstoragechain";

export interface MsgCreateResource {
  creator: string;
  index: string;
  cid: string;
}

export interface MsgCreateResourceResponse {
}

export interface MsgUpdateResource {
  creator: string;
  index: string;
  cid: string;
}

export interface MsgUpdateResourceResponse {
}

export interface MsgDeleteResource {
  creator: string;
  index: string;
}

export interface MsgDeleteResourceResponse {
}

function createBaseMsgCreateResource(): MsgCreateResource {
  return { creator: "", index: "", cid: "" };
}

export const MsgCreateResource = {
  encode(message: MsgCreateResource, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    if (message.cid !== "") {
      writer.uint32(26).string(message.cid);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateResource {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateResource();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        case 3:
          message.cid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateResource {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      index: isSet(object.index) ? String(object.index) : "",
      cid: isSet(object.cid) ? String(object.cid) : "",
    };
  },

  toJSON(message: MsgCreateResource): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    message.cid !== undefined && (obj.cid = message.cid);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateResource>, I>>(object: I): MsgCreateResource {
    const message = createBaseMsgCreateResource();
    message.creator = object.creator ?? "";
    message.index = object.index ?? "";
    message.cid = object.cid ?? "";
    return message;
  },
};

function createBaseMsgCreateResourceResponse(): MsgCreateResourceResponse {
  return {};
}

export const MsgCreateResourceResponse = {
  encode(_: MsgCreateResourceResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateResourceResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateResourceResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateResourceResponse {
    return {};
  },

  toJSON(_: MsgCreateResourceResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateResourceResponse>, I>>(_: I): MsgCreateResourceResponse {
    const message = createBaseMsgCreateResourceResponse();
    return message;
  },
};

function createBaseMsgUpdateResource(): MsgUpdateResource {
  return { creator: "", index: "", cid: "" };
}

export const MsgUpdateResource = {
  encode(message: MsgUpdateResource, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    if (message.cid !== "") {
      writer.uint32(26).string(message.cid);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateResource {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateResource();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        case 3:
          message.cid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateResource {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      index: isSet(object.index) ? String(object.index) : "",
      cid: isSet(object.cid) ? String(object.cid) : "",
    };
  },

  toJSON(message: MsgUpdateResource): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    message.cid !== undefined && (obj.cid = message.cid);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateResource>, I>>(object: I): MsgUpdateResource {
    const message = createBaseMsgUpdateResource();
    message.creator = object.creator ?? "";
    message.index = object.index ?? "";
    message.cid = object.cid ?? "";
    return message;
  },
};

function createBaseMsgUpdateResourceResponse(): MsgUpdateResourceResponse {
  return {};
}

export const MsgUpdateResourceResponse = {
  encode(_: MsgUpdateResourceResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateResourceResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateResourceResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateResourceResponse {
    return {};
  },

  toJSON(_: MsgUpdateResourceResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateResourceResponse>, I>>(_: I): MsgUpdateResourceResponse {
    const message = createBaseMsgUpdateResourceResponse();
    return message;
  },
};

function createBaseMsgDeleteResource(): MsgDeleteResource {
  return { creator: "", index: "" };
}

export const MsgDeleteResource = {
  encode(message: MsgDeleteResource, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteResource {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteResource();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteResource {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      index: isSet(object.index) ? String(object.index) : "",
    };
  },

  toJSON(message: MsgDeleteResource): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteResource>, I>>(object: I): MsgDeleteResource {
    const message = createBaseMsgDeleteResource();
    message.creator = object.creator ?? "";
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseMsgDeleteResourceResponse(): MsgDeleteResourceResponse {
  return {};
}

export const MsgDeleteResourceResponse = {
  encode(_: MsgDeleteResourceResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteResourceResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteResourceResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDeleteResourceResponse {
    return {};
  },

  toJSON(_: MsgDeleteResourceResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteResourceResponse>, I>>(_: I): MsgDeleteResourceResponse {
    const message = createBaseMsgDeleteResourceResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateResource(request: MsgCreateResource): Promise<MsgCreateResourceResponse>;
  UpdateResource(request: MsgUpdateResource): Promise<MsgUpdateResourceResponse>;
  DeleteResource(request: MsgDeleteResource): Promise<MsgDeleteResourceResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateResource = this.CreateResource.bind(this);
    this.UpdateResource = this.UpdateResource.bind(this);
    this.DeleteResource = this.DeleteResource.bind(this);
  }
  CreateResource(request: MsgCreateResource): Promise<MsgCreateResourceResponse> {
    const data = MsgCreateResource.encode(request).finish();
    const promise = this.rpc.request("blockcraftsman.warmstoragechain.warmstoragechain.Msg", "CreateResource", data);
    return promise.then((data) => MsgCreateResourceResponse.decode(new _m0.Reader(data)));
  }

  UpdateResource(request: MsgUpdateResource): Promise<MsgUpdateResourceResponse> {
    const data = MsgUpdateResource.encode(request).finish();
    const promise = this.rpc.request("blockcraftsman.warmstoragechain.warmstoragechain.Msg", "UpdateResource", data);
    return promise.then((data) => MsgUpdateResourceResponse.decode(new _m0.Reader(data)));
  }

  DeleteResource(request: MsgDeleteResource): Promise<MsgDeleteResourceResponse> {
    const data = MsgDeleteResource.encode(request).finish();
    const promise = this.rpc.request("blockcraftsman.warmstoragechain.warmstoragechain.Msg", "DeleteResource", data);
    return promise.then((data) => MsgDeleteResourceResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
