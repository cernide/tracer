// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.1.5
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface V1WasbType
 */
export interface V1WasbType {
    /**
     * 
     * @type {string}
     * @memberof V1WasbType
     */
    container?: string;
    /**
     * 
     * @type {string}
     * @memberof V1WasbType
     */
    storage_account?: string;
    /**
     * 
     * @type {boolean}
     * @memberof V1WasbType
     */
    path?: boolean;
}

export function V1WasbTypeFromJSON(json: any): V1WasbType {
    return V1WasbTypeFromJSONTyped(json, false);
}

export function V1WasbTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1WasbType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'container': !exists(json, 'container') ? undefined : json['container'],
        'storage_account': !exists(json, 'storage_account') ? undefined : json['storage_account'],
        'path': !exists(json, 'path') ? undefined : json['path'],
    };
}

export function V1WasbTypeToJSON(value?: V1WasbType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'container': value.container,
        'storage_account': value.storage_account,
        'path': value.path,
    };
}


