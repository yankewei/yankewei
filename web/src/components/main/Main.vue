<template>
    <div class="w-full -mt-14 pl-24 pr-10 h-auto flex">
        <div class="w-3/5">
            <Content/>
        </div>
        <div class="w-2/5">
            <Sidebar/>
        </div>
    </div>
</template>


<script>

import Content from "./content/Content.vue"
import Sidebar from "./Sidebar.vue"

export default {
    name: 'Main',
    components: {
        Content,
        Sidebar
    }
}
</script>