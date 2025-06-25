/**
 * @param {*} obj
 * @param {*} classFunction
 * @return {boolean}
 */
var checkIfInstanceOf = function(obj, classFunction) {
  if(obj===null||classFunction===null||typeof classFunction!=='function') {
    return false;
  }
  if(obj instanceof classFunction) {
    return true;
  }

};