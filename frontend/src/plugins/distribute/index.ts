import Constant from "@/global/Constant";
import PluginModeEnum from "@/enumeration/PluginModeEnum";
import {web, webPreload, DbStorage} from "@/plugins/distribute/web";
import {tauri, tauriPreload} from "@/plugins/distribute/tauri";
import {serverPreload} from "@/plugins/distribute/server";

let utools = web;
let source= webPreload;


export const instance = utools;
// @ts-ignore
export const preload = source;
// @ts-ignore

export const dbStorage = window.utools ? window.utools.dbStorage : DbStorage
