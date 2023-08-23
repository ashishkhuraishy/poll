/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../poll/params";
import { PollInfo } from "../poll/poll_info";

export const protobufPackage = "ashishkhuraishy.poll.poll";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetPollInfoRequest {}

export interface QueryGetPollInfoResponse {
  PollInfo: PollInfo | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
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

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetPollInfoRequest: object = {};

export const QueryGetPollInfoRequest = {
  encode(_: QueryGetPollInfoRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPollInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetPollInfoRequest,
    } as QueryGetPollInfoRequest;
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

  fromJSON(_: any): QueryGetPollInfoRequest {
    const message = {
      ...baseQueryGetPollInfoRequest,
    } as QueryGetPollInfoRequest;
    return message;
  },

  toJSON(_: QueryGetPollInfoRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetPollInfoRequest>
  ): QueryGetPollInfoRequest {
    const message = {
      ...baseQueryGetPollInfoRequest,
    } as QueryGetPollInfoRequest;
    return message;
  },
};

const baseQueryGetPollInfoResponse: object = {};

export const QueryGetPollInfoResponse = {
  encode(
    message: QueryGetPollInfoResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.PollInfo !== undefined) {
      PollInfo.encode(message.PollInfo, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetPollInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetPollInfoResponse,
    } as QueryGetPollInfoResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.PollInfo = PollInfo.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetPollInfoResponse {
    const message = {
      ...baseQueryGetPollInfoResponse,
    } as QueryGetPollInfoResponse;
    if (object.PollInfo !== undefined && object.PollInfo !== null) {
      message.PollInfo = PollInfo.fromJSON(object.PollInfo);
    } else {
      message.PollInfo = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetPollInfoResponse): unknown {
    const obj: any = {};
    message.PollInfo !== undefined &&
      (obj.PollInfo = message.PollInfo
        ? PollInfo.toJSON(message.PollInfo)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetPollInfoResponse>
  ): QueryGetPollInfoResponse {
    const message = {
      ...baseQueryGetPollInfoResponse,
    } as QueryGetPollInfoResponse;
    if (object.PollInfo !== undefined && object.PollInfo !== null) {
      message.PollInfo = PollInfo.fromPartial(object.PollInfo);
    } else {
      message.PollInfo = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a PollInfo by index. */
  PollInfo(request: QueryGetPollInfoRequest): Promise<QueryGetPollInfoResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "ashishkhuraishy.poll.poll.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  PollInfo(
    request: QueryGetPollInfoRequest
  ): Promise<QueryGetPollInfoResponse> {
    const data = QueryGetPollInfoRequest.encode(request).finish();
    const promise = this.rpc.request(
      "ashishkhuraishy.poll.poll.Query",
      "PollInfo",
      data
    );
    return promise.then((data) =>
      QueryGetPollInfoResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
