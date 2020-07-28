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
 * @interface V1Installation
 */
export interface V1Installation {
    /**
     * 
     * @type {string}
     * @memberof V1Installation
     */
    key?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Installation
     */
    version?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Installation
     */
    dist?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof V1Installation
     */
    auth?: Array<string>;
}

export function V1InstallationFromJSON(json: any): V1Installation {
    return V1InstallationFromJSONTyped(json, false);
}

export function V1InstallationFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1Installation {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'key': !exists(json, 'key') ? undefined : json['key'],
        'version': !exists(json, 'version') ? undefined : json['version'],
        'dist': !exists(json, 'dist') ? undefined : json['dist'],
        'auth': !exists(json, 'auth') ? undefined : json['auth'],
    };
}

export function V1InstallationToJSON(value?: V1Installation | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'key': value.key,
        'version': value.version,
        'dist': value.dist,
        'auth': value.auth,
    };
}


