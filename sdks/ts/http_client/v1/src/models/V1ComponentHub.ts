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
 * @interface V1ComponentHub
 */
export interface V1ComponentHub {
    /**
     * 
     * @type {string}
     * @memberof V1ComponentHub
     */
    uuid?: string;
    /**
     * 
     * @type {string}
     * @memberof V1ComponentHub
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof V1ComponentHub
     */
    content?: string;
    /**
     * 
     * @type {string}
     * @memberof V1ComponentHub
     */
    description?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof V1ComponentHub
     */
    tags?: Array<string>;
    /**
     * 
     * @type {boolean}
     * @memberof V1ComponentHub
     */
    disabled?: boolean;
    /**
     * 
     * @type {boolean}
     * @memberof V1ComponentHub
     */
    deleted?: boolean;
    /**
     * 
     * @type {Date}
     * @memberof V1ComponentHub
     */
    created_at?: Date;
    /**
     * 
     * @type {Date}
     * @memberof V1ComponentHub
     */
    updated_at?: Date;
}

export function V1ComponentHubFromJSON(json: any): V1ComponentHub {
    return V1ComponentHubFromJSONTyped(json, false);
}

export function V1ComponentHubFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1ComponentHub {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'uuid': !exists(json, 'uuid') ? undefined : json['uuid'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'content': !exists(json, 'content') ? undefined : json['content'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'tags': !exists(json, 'tags') ? undefined : json['tags'],
        'disabled': !exists(json, 'disabled') ? undefined : json['disabled'],
        'deleted': !exists(json, 'deleted') ? undefined : json['deleted'],
        'created_at': !exists(json, 'created_at') ? undefined : (new Date(json['created_at'])),
        'updated_at': !exists(json, 'updated_at') ? undefined : (new Date(json['updated_at'])),
    };
}

export function V1ComponentHubToJSON(value?: V1ComponentHub | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'uuid': value.uuid,
        'name': value.name,
        'content': value.content,
        'description': value.description,
        'tags': value.tags,
        'disabled': value.disabled,
        'deleted': value.deleted,
        'created_at': value.created_at === undefined ? undefined : (value.created_at.toISOString()),
        'updated_at': value.updated_at === undefined ? undefined : (value.updated_at.toISOString()),
    };
}


