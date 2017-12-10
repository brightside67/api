import dispatcher from 'edispatcher'
import domd from 'domd'


const d = domd(document.body)

const componentMap = {}

const element = document.getElementsByTagName('body')[0]
const components = element.querySelectorAll('[data-component]')

Array.prototype.forEach.call(components, (mount_point) => {
  const cn = mount_point.getAttribute('data-component')
  try {
    const ck = cn + ':' + (mount_point.getAttribute('data-component-id') || 'default')
    // If we should call already inited destructor
    if (componentMap[ck] && typeof(componentMap[ck]) === 'function') {
      componentMap[ck]()
    }
    const cc = require('./component/' + cn + '/index.js').default
    componentMap[ck] = new cc(mount_point) // Get component destructor
  } catch (e) {
    console.warn('Failed to initialize component ' + cn + '. ' + e)
  }
})