cmake_minimum_required(VERSION 3.11)

option(WASI_SDK_PREFIX "Path to WASI SDK")

add_compile_options(-Os -DNDEBUG -flto)
add_link_options(-s -flto)

set(CMAKE_MODULE_PATH ${CMAKE_SOURCE_DIR}/cmake)
set(CMAKE_TOOLCHAIN_FILE "${WASI_SDK_PREFIX}/share/cmake/wasi-sdk.cmake")

project(go-graphviz)

include(FetchContent)

FetchContent_Declare(
    graphviz
    URL https://gitlab.com/api/v4/projects/4207231/packages/generic/graphviz-releases/7.0.6/graphviz-7.0.6.tar.gz
    URL_HASH SHA256=8f979c23323c3e877f824dd1bd204de691c98d31cd3e7c507b693f1adea9f583
    DOWNLOAD_EXTRACT_TIMESTAMP FALSE
    PATCH_COMMAND patch -p1 < ${CMAKE_CURRENT_LIST_DIR}/patch0
)

FetchContent_MakeAvailable(graphviz)

add_executable(dot
    ${graphviz_SOURCE_DIR}/cmd/dot/dot.c
    ${graphviz_SOURCE_DIR}/cmd/dot/dot_builtins.cpp
    ${graphviz_SOURCE_DIR}/lib/common/htmlparse.c
    ${graphviz_SOURCE_DIR}/lib/cgraph/grammar.c
    ${graphviz_SOURCE_DIR}/lib/cgraph/scan.c
    ${graphviz_SOURCE_DIR}/lib/cgraph/pend.c
    ${graphviz_SOURCE_DIR}/lib/cgraph/edge.c
    ${graphviz_SOURCE_DIR}/lib/cgraph/write.c
    ${graphviz_SOURCE_DIR}/lib/cgraph/flatten.c
    ${graphviz_SOURCE_DIR}/lib/cgraph/agerror.c
    ${graphviz_SOURCE_DIR}/lib/cgraph/id.c
    ${graphviz_SOURCE_DIR}/lib/cgraph/graph.c
    ${graphviz_SOURCE_DIR}/lib/cgraph/refstr.c
    ${graphviz_SOURCE_DIR}/lib/cgraph/rec.c
    ${graphviz_SOURCE_DIR}/lib/cgraph/subg.c
    ${graphviz_SOURCE_DIR}/lib/cgraph/attr.c
    ${graphviz_SOURCE_DIR}/lib/cgraph/io.c
    ${graphviz_SOURCE_DIR}/lib/cgraph/apply.c
    ${graphviz_SOURCE_DIR}/lib/cgraph/obj.c
    ${graphviz_SOURCE_DIR}/lib/cgraph/mem.c
    ${graphviz_SOURCE_DIR}/lib/cgraph/imap.c
    ${graphviz_SOURCE_DIR}/lib/cgraph/node.c
    ${graphviz_SOURCE_DIR}/lib/cgraph/utils.c
    ${graphviz_SOURCE_DIR}/lib/cdt/dtwalk.c
    ${graphviz_SOURCE_DIR}/lib/cdt/dtrestore.c
    ${graphviz_SOURCE_DIR}/lib/cdt/dtlist.c
    ${graphviz_SOURCE_DIR}/lib/cdt/dttree.c
    ${graphviz_SOURCE_DIR}/lib/cdt/dtsize.c
    ${graphviz_SOURCE_DIR}/lib/cdt/dthash.c
    ${graphviz_SOURCE_DIR}/lib/cdt/dtstat.c
    ${graphviz_SOURCE_DIR}/lib/cdt/dtview.c
    ${graphviz_SOURCE_DIR}/lib/cdt/dtdisc.c
    ${graphviz_SOURCE_DIR}/lib/cdt/dtextract.c
    ${graphviz_SOURCE_DIR}/lib/cdt/dtmethod.c
    ${graphviz_SOURCE_DIR}/lib/cdt/dtflatten.c
    ${graphviz_SOURCE_DIR}/lib/cdt/dtclose.c
    ${graphviz_SOURCE_DIR}/lib/cdt/dtrenew.c
    ${graphviz_SOURCE_DIR}/lib/cdt/dtstrhash.c
    ${graphviz_SOURCE_DIR}/lib/cdt/dtopen.c
    ${graphviz_SOURCE_DIR}/lib/pack/pack.c
    ${graphviz_SOURCE_DIR}/lib/pack/ccomps.c
    ${graphviz_SOURCE_DIR}/lib/rbtree/stack.c
    ${graphviz_SOURCE_DIR}/lib/rbtree/red_black_tree.c
    ${graphviz_SOURCE_DIR}/lib/rbtree/misc.c
    ${graphviz_SOURCE_DIR}/lib/expr/extoken.c
    ${graphviz_SOURCE_DIR}/lib/expr/exerror.c
    ${graphviz_SOURCE_DIR}/lib/expr/exnospace.c
    ${graphviz_SOURCE_DIR}/lib/expr/exdata.c
    ${graphviz_SOURCE_DIR}/lib/expr/exstash.c
    ${graphviz_SOURCE_DIR}/lib/expr/excc.c
    ${graphviz_SOURCE_DIR}/lib/expr/exexpr.c
    ${graphviz_SOURCE_DIR}/lib/expr/exzero.c
    ${graphviz_SOURCE_DIR}/lib/expr/exeval.c
    ${graphviz_SOURCE_DIR}/lib/expr/extype.c
    ${graphviz_SOURCE_DIR}/lib/expr/exopen.c
    ${graphviz_SOURCE_DIR}/lib/expr/excontext.c
    ${graphviz_SOURCE_DIR}/lib/fdpgen/xlayout.c
    ${graphviz_SOURCE_DIR}/lib/fdpgen/grid.c
    ${graphviz_SOURCE_DIR}/lib/fdpgen/fdpinit.c
    ${graphviz_SOURCE_DIR}/lib/fdpgen/comp.c
    ${graphviz_SOURCE_DIR}/lib/fdpgen/layout.c
    ${graphviz_SOURCE_DIR}/lib/fdpgen/clusteredges.c
    ${graphviz_SOURCE_DIR}/lib/fdpgen/tlayout.c
    ${graphviz_SOURCE_DIR}/lib/fdpgen/dbg.c
    ${graphviz_SOURCE_DIR}/lib/osage/osageinit.c
    ${graphviz_SOURCE_DIR}/lib/ortho/maze.c
    ${graphviz_SOURCE_DIR}/lib/ortho/sgraph.c
    ${graphviz_SOURCE_DIR}/lib/ortho/partition.c
    ${graphviz_SOURCE_DIR}/lib/ortho/trapezoid.c
    ${graphviz_SOURCE_DIR}/lib/ortho/rawgraph.c
    ${graphviz_SOURCE_DIR}/lib/ortho/fPQ.c
    ${graphviz_SOURCE_DIR}/lib/ortho/ortho.c
    ${graphviz_SOURCE_DIR}/lib/vmalloc/vmclear.c
    ${graphviz_SOURCE_DIR}/lib/vmalloc/vmstrdup.c
    ${graphviz_SOURCE_DIR}/lib/vmalloc/vmclose.c
    ${graphviz_SOURCE_DIR}/lib/vmalloc/vmalloc.c
    ${graphviz_SOURCE_DIR}/lib/vmalloc/vmopen.c
    ${graphviz_SOURCE_DIR}/lib/pathplan/route.c
    ${graphviz_SOURCE_DIR}/lib/pathplan/inpoly.c
    ${graphviz_SOURCE_DIR}/lib/pathplan/visibility.c
    ${graphviz_SOURCE_DIR}/lib/pathplan/shortestpth.c
    ${graphviz_SOURCE_DIR}/lib/pathplan/util.c
    ${graphviz_SOURCE_DIR}/lib/pathplan/cvt.c
    ${graphviz_SOURCE_DIR}/lib/pathplan/shortest.c
    ${graphviz_SOURCE_DIR}/lib/pathplan/triang.c
    ${graphviz_SOURCE_DIR}/lib/pathplan/solvers.c
    ${graphviz_SOURCE_DIR}/lib/edgepaint/node_distinct_coloring.c
    ${graphviz_SOURCE_DIR}/lib/edgepaint/edge_distinct_coloring.c
    ${graphviz_SOURCE_DIR}/lib/edgepaint/intersection.c
    ${graphviz_SOURCE_DIR}/lib/edgepaint/lab_gamut.c
    ${graphviz_SOURCE_DIR}/lib/edgepaint/lab.c
    ${graphviz_SOURCE_DIR}/lib/edgepaint/furtherest_point.c
    ${graphviz_SOURCE_DIR}/lib/label/xlabels.c
    ${graphviz_SOURCE_DIR}/lib/label/rectangle.c
    ${graphviz_SOURCE_DIR}/lib/label/split.q.c
    ${graphviz_SOURCE_DIR}/lib/label/node.c
    ${graphviz_SOURCE_DIR}/lib/label/index.c
    ${graphviz_SOURCE_DIR}/lib/ast/pathaccess.c
    ${graphviz_SOURCE_DIR}/lib/ast/chrtoi.c
    ${graphviz_SOURCE_DIR}/lib/ast/fmtbuf.c
    ${graphviz_SOURCE_DIR}/lib/ast/pathcat.c
    ${graphviz_SOURCE_DIR}/lib/ast/strcopy.c
    ${graphviz_SOURCE_DIR}/lib/ast/strmatch.c
    ${graphviz_SOURCE_DIR}/lib/ast/pathpath.c
    ${graphviz_SOURCE_DIR}/lib/ast/stresc.c
    ${graphviz_SOURCE_DIR}/lib/ast/pathcanon.c
    ${graphviz_SOURCE_DIR}/lib/ast/chresc.c
    ${graphviz_SOURCE_DIR}/lib/ast/fmtesc.c
    ${graphviz_SOURCE_DIR}/lib/ast/error.c
    ${graphviz_SOURCE_DIR}/lib/ast/pathfind.c
    ${graphviz_SOURCE_DIR}/lib/patchwork/tree_map.c
    ${graphviz_SOURCE_DIR}/lib/patchwork/patchwork.c
    ${graphviz_SOURCE_DIR}/lib/patchwork/patchworkinit.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/edges.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/solve.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/lu.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/printvis.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/randomkit.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/poly.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/neatoinit.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/pca.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/geometry.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/neatosplines.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/multispline.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/embed_graph.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/constrained_majorization.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/sgd.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/matinv.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/site.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/adjust.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/info.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/heap.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/stuff.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/voronoi.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/dijkstra.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/bfs.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/compute_hierarchy.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/constraint.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/overlap.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/quad_prog_vpsc.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/circuit.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/matrix_ops.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/hedges.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/smart_ini_x.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/call_tri.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/closest.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/constrained_majorization_ipsep.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/quad_prog_solve.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/conjgrad.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/legal.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/stress.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/opt_arrangement.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/memory.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/delaunay.c
    ${graphviz_SOURCE_DIR}/lib/neatogen/kkutils.c
    ${graphviz_SOURCE_DIR}/lib/twopigen/twopiinit.c
    ${graphviz_SOURCE_DIR}/lib/twopigen/circle.c
    ${graphviz_SOURCE_DIR}/lib/dotgen/flat.c
    ${graphviz_SOURCE_DIR}/lib/dotgen/aspect.c
    ${graphviz_SOURCE_DIR}/lib/dotgen/class2.c
    ${graphviz_SOURCE_DIR}/lib/dotgen/decomp.c
    ${graphviz_SOURCE_DIR}/lib/dotgen/cluster.c
    ${graphviz_SOURCE_DIR}/lib/dotgen/dotinit.c
    ${graphviz_SOURCE_DIR}/lib/dotgen/fastgr.c
    ${graphviz_SOURCE_DIR}/lib/dotgen/acyclic.c
    ${graphviz_SOURCE_DIR}/lib/dotgen/conc.c
    ${graphviz_SOURCE_DIR}/lib/dotgen/position.c
    ${graphviz_SOURCE_DIR}/lib/dotgen/rank.c
    ${graphviz_SOURCE_DIR}/lib/dotgen/sameport.c
    ${graphviz_SOURCE_DIR}/lib/dotgen/compound.c
    ${graphviz_SOURCE_DIR}/lib/dotgen/dotsplines.c
    ${graphviz_SOURCE_DIR}/lib/dotgen/class1.c
    ${graphviz_SOURCE_DIR}/lib/dotgen/mincross.c
    ${graphviz_SOURCE_DIR}/lib/sparse/colorutil.c
    ${graphviz_SOURCE_DIR}/lib/sparse/SparseMatrix.c
    ${graphviz_SOURCE_DIR}/lib/sparse/DotIO.c
    ${graphviz_SOURCE_DIR}/lib/sparse/mq.c
    ${graphviz_SOURCE_DIR}/lib/sparse/color_palette.c
    ${graphviz_SOURCE_DIR}/lib/sparse/clustering.c
    ${graphviz_SOURCE_DIR}/lib/sparse/BinaryHeap.c
    ${graphviz_SOURCE_DIR}/lib/sparse/LinkedList.c
    ${graphviz_SOURCE_DIR}/lib/sparse/QuadTree.c
    ${graphviz_SOURCE_DIR}/lib/sparse/general.c
    ${graphviz_SOURCE_DIR}/lib/ingraphs/ingraphs.c
    ${graphviz_SOURCE_DIR}/lib/sfio/sfvscanf.c
    ${graphviz_SOURCE_DIR}/lib/sfio/sfwrite.c
    ${graphviz_SOURCE_DIR}/lib/sfio/sfexcept.c
    ${graphviz_SOURCE_DIR}/lib/sfio/sffilbuf.c
    ${graphviz_SOURCE_DIR}/lib/sfio/sfscanf.c
    ${graphviz_SOURCE_DIR}/lib/sfio/sftable.c
    ${graphviz_SOURCE_DIR}/lib/sfio/sfwr.c
    ${graphviz_SOURCE_DIR}/lib/sfio/sfsk.c
    ${graphviz_SOURCE_DIR}/lib/sfio/sfmode.c
    ${graphviz_SOURCE_DIR}/lib/sfio/sfsync.c
    ${graphviz_SOURCE_DIR}/lib/sfio/sfstack.c
    ${graphviz_SOURCE_DIR}/lib/sfio/sfflsbuf.c
    ${graphviz_SOURCE_DIR}/lib/sfio/sfrd.c
    ${graphviz_SOURCE_DIR}/lib/sfio/sfputr.c
    ${graphviz_SOURCE_DIR}/lib/sfio/sfnew.c
    ${graphviz_SOURCE_DIR}/lib/sfio/sfungetc.c
    ${graphviz_SOURCE_DIR}/lib/sfio/sfcvt.c
    ${graphviz_SOURCE_DIR}/lib/sfio/sfdisc.c
    ${graphviz_SOURCE_DIR}/lib/sfio/sfseek.c
    ${graphviz_SOURCE_DIR}/lib/sfio/sfopen.c
    ${graphviz_SOURCE_DIR}/lib/sfio/sfvprintf.c
    ${graphviz_SOURCE_DIR}/lib/sfio/sfnputc.c
    ${graphviz_SOURCE_DIR}/lib/sfio/sfclose.c
    ${graphviz_SOURCE_DIR}/lib/sfio/sfread.c
    ${graphviz_SOURCE_DIR}/lib/sfio/sfpkrd.c
    ${graphviz_SOURCE_DIR}/lib/sfio/sfextern.c
    ${graphviz_SOURCE_DIR}/lib/sfio/sfprintf.c
    ${graphviz_SOURCE_DIR}/lib/sfio/sfsetfd.c
    ${graphviz_SOURCE_DIR}/lib/sfio/Sfio_f/_sffileno.c
    ${graphviz_SOURCE_DIR}/lib/sfio/Sfio_f/_sfputc.c
    ${graphviz_SOURCE_DIR}/lib/sfio/Sfio_f/_sfslen.c
    ${graphviz_SOURCE_DIR}/lib/sfio/Sfio_f/_sfgetc.c
    ${graphviz_SOURCE_DIR}/lib/sfio/sfraise.c
    ${graphviz_SOURCE_DIR}/lib/sfio/sfswap.c
    ${graphviz_SOURCE_DIR}/lib/xdot/xdot.c
    ${graphviz_SOURCE_DIR}/lib/common/postproc.c
    ${graphviz_SOURCE_DIR}/lib/common/output.c
    ${graphviz_SOURCE_DIR}/lib/common/pointset.c
    ${graphviz_SOURCE_DIR}/lib/common/labels.c
    ${graphviz_SOURCE_DIR}/lib/common/geom.c
    ${graphviz_SOURCE_DIR}/lib/common/arrows.c
    ${graphviz_SOURCE_DIR}/lib/common/intset.c
    ${graphviz_SOURCE_DIR}/lib/common/htmllex.c
    ${graphviz_SOURCE_DIR}/lib/common/ellipse.c
    ${graphviz_SOURCE_DIR}/lib/common/emit.c
    ${graphviz_SOURCE_DIR}/lib/common/args.c
    ${graphviz_SOURCE_DIR}/lib/common/splines.c
    ${graphviz_SOURCE_DIR}/lib/common/psusershape.c
    ${graphviz_SOURCE_DIR}/lib/common/htmltable.c
    ${graphviz_SOURCE_DIR}/lib/common/input.c
    ${graphviz_SOURCE_DIR}/lib/common/ns.c
    ${graphviz_SOURCE_DIR}/lib/common/textspan.c
    ${graphviz_SOURCE_DIR}/lib/common/globals.c
    ${graphviz_SOURCE_DIR}/lib/common/colxlate.c
    ${graphviz_SOURCE_DIR}/lib/common/taper.c
    ${graphviz_SOURCE_DIR}/lib/common/routespl.c
    ${graphviz_SOURCE_DIR}/lib/common/textspan_lut.c
    ${graphviz_SOURCE_DIR}/lib/common/xml.c
    ${graphviz_SOURCE_DIR}/lib/common/shapes.c
    ${graphviz_SOURCE_DIR}/lib/common/utils.c
    ${graphviz_SOURCE_DIR}/lib/common/memory.c
    ${graphviz_SOURCE_DIR}/lib/common/timing.c
    ${graphviz_SOURCE_DIR}/lib/gvc/gvtextlayout.c
    ${graphviz_SOURCE_DIR}/lib/gvc/gvplugin.c
    ${graphviz_SOURCE_DIR}/lib/gvc/gvrender.c
    ${graphviz_SOURCE_DIR}/lib/gvc/gvjobs.c
    ${graphviz_SOURCE_DIR}/lib/gvc/gvevent.c
    ${graphviz_SOURCE_DIR}/lib/gvc/gvtool_tred.c
    ${graphviz_SOURCE_DIR}/lib/gvc/gvusershape.c
    ${graphviz_SOURCE_DIR}/lib/gvc/gvloadimage.c
    ${graphviz_SOURCE_DIR}/lib/gvc/gvconfig.c
    ${graphviz_SOURCE_DIR}/lib/gvc/gvlayout.c
    ${graphviz_SOURCE_DIR}/lib/gvc/gvc.c
    ${graphviz_SOURCE_DIR}/lib/gvc/gvdevice.c
    ${graphviz_SOURCE_DIR}/lib/gvc/gvcontext.c
    ${graphviz_SOURCE_DIR}/lib/circogen/block.c
    ${graphviz_SOURCE_DIR}/lib/circogen/blocktree.c
    ${graphviz_SOURCE_DIR}/lib/circogen/circpos.c
    ${graphviz_SOURCE_DIR}/lib/circogen/circular.c
    ${graphviz_SOURCE_DIR}/lib/circogen/blockpath.c
    ${graphviz_SOURCE_DIR}/lib/circogen/circularinit.c
    ${graphviz_SOURCE_DIR}/lib/circogen/edgelist.c
    ${graphviz_SOURCE_DIR}/lib/circogen/nodelist.c
    ${graphviz_SOURCE_DIR}/plugin/core/gvrender_core_map.c
    ${graphviz_SOURCE_DIR}/plugin/core/gvrender_core_tk.c
    ${graphviz_SOURCE_DIR}/plugin/core/gvrender_core_svg.c
    ${graphviz_SOURCE_DIR}/plugin/core/gvrender_core_ps.c
    ${graphviz_SOURCE_DIR}/plugin/core/gvrender_core_vml.c
    ${graphviz_SOURCE_DIR}/plugin/core/gvplugin_core.c
    ${graphviz_SOURCE_DIR}/plugin/core/gvrender_core_pic.c
    ${graphviz_SOURCE_DIR}/plugin/core/gvrender_core_fig.c
    ${graphviz_SOURCE_DIR}/plugin/core/gvloadimage_core.c
    ${graphviz_SOURCE_DIR}/plugin/core/gvrender_core_pov.c
    ${graphviz_SOURCE_DIR}/plugin/core/gvrender_core_json.c
    ${graphviz_SOURCE_DIR}/plugin/core/gvrender_core_mp.c
    ${graphviz_SOURCE_DIR}/plugin/core/gvrender_core_dot.c
    ${graphviz_SOURCE_DIR}/plugin/neato_layout/gvplugin_neato_layout.c
    ${graphviz_SOURCE_DIR}/plugin/neato_layout/gvlayout_neato_layout.c
    ${graphviz_SOURCE_DIR}/plugin/dot_layout/gvlayout_dot_layout.c
    ${graphviz_SOURCE_DIR}/plugin/dot_layout/gvplugin_dot_layout.c
    ${graphviz_SOURCE_DIR}/tmpfile.c
)
target_include_directories(dot PRIVATE
    ${graphviz_SOURCE_DIR}/
    ${graphviz_SOURCE_DIR}/lib/
    ${graphviz_SOURCE_DIR}/lib/cdt/
    ${graphviz_SOURCE_DIR}/lib/cgraph/
    ${graphviz_SOURCE_DIR}/lib/common/
    ${graphviz_SOURCE_DIR}/lib/gvc/
    ${graphviz_SOURCE_DIR}/lib/gvpr/
    ${graphviz_SOURCE_DIR}/lib/pathplan/
)
target_compile_definitions(dot PRIVATE
    DEMAND_LOADING=0
    _WASI_EMULATED_SIGNAL
    _WASI_EMULATED_PROCESS_CLOCKS
    _WASI_EMULATED_GETPID
)
target_link_libraries(dot PRIVATE
    wasi-emulated-signal
    wasi-emulated-process-clocks
    wasi-emulated-getpid
)