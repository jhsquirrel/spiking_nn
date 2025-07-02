package main
import (
    "fmt"
    "math"
    //"os"
    //"strconv"
)

type Synapse struct {
    weight float64
    startnode int
    endnode int
}

type Neuron struct {
    a float64
    b float64
    c float64
    d float64
    u float64
    v float64
    fired bool
    I float64
    I_tick float64
    v_tick float64
    nodeid int
}

//func neuron_iz(a float64, b float64, c float64, d float64, I float64, u float64, v float64) (float64, float64, bool) {
func neuron_iz(n Neuron, I float64) (float64, float64, float64, bool) {
    //var u_tick float64
    var v_tick float64
    //var v_tick2 float64
    var a float64 = n.a
    var b float64 = n.b
    var c float64 = n.c
    var d float64 = n.d
    //var I float64 = n.I
    var u float64 = n.u
    var v float64 = n.v
    var fired bool = false
    I = I + n.I_tick

    //v_tick = (( 0.04 * math.Pow(v, 2) ) + ( 5 * v ) + 140.0 - u + I)
    t1 := 0.04 * math.Pow(v, 2)
    //fmt.Println("#t1", t1, v)
    t2 := 5.0 * v
    //fmt.Println("#t2", t2, v)
    t3 := 140.0 - u + I
    //fmt.Println("#t3", t3, u, I)
    //v = v + 0.02 * (t1 + t2 + t3)
    v = v + (t1 + t2 + t3)
    //v_tick2 = v_tick + 0.5 * ( 0.04 * math.Pow(v_tick, 2) ) + ( 5 * v_tick ) + 140.0 - u + I
    
    //u = u + 0.02 * a * (  b * v - u )
    u = u + a * (  b * v - u )

    //fmt.Println("#", t1, t2, t3, "u=", u, "v=", v, I)

    //v = v_tick
    //u = u_tick
    v_tick = v

    if v > 30.0 {
        v_tick = v   
        v = c
        u = u + d
        fired = true
    }

    return u, v, v_tick, fired
}

func makeCurrent(maxt int, start int, end int, Ival float64) ([] float64) {
    var I[] float64
    for i := 0; i < maxt; i++ {
        if i > start && i < end {
           I = append(I, Ival)
        } else {
           I = append(I, 0.0)
        }
    }
    return I 
}

func makeNeurons(total int, startid int, net[] Neuron, a float64, b float64, c float64, d float64) ([] Neuron) {
    // initial value of v
    var v = c
    // initial value of u
    var u = b * v
    for i := 0; i < total; i++ {
        //fmt.Println(a, b, c, d, u, v, i)
        net = append(net, Neuron{a, b, c, d, u, v, false, 0, 0, 0, startid + i})
    }
    return net
}
func makeRSNeurons(total int, startid int, net[] Neuron) ([] Neuron) {
    var a, b, c, d float64
    a = 0.02
    b = 0.2
    c = -65
    d = 8
    net = makeNeurons(total, startid, net, a, b, c, d)
    return net
}

func makeIBNeurons(total int, startid int, net[] Neuron) ([] Neuron) {
    var a, b, c, d float64
    a = 0.02
    b = 0.2
    c = -55
    d = 4
    net = makeNeurons(total, startid, net, a, b, c, d)
    return net
}

func makeCHNeurons(total int, startid int, net[] Neuron) ([] Neuron) {
    var a, b, c, d float64
    a = 0.02
    b = 0.2
    c = -50
    d = 2
    net = makeNeurons(total, startid, net, a, b, c, d)
    return net
}

func makeFSNeurons(total int, startid int, net[] Neuron) ([] Neuron) {
    var a, b, c, d float64
    a = 0.1
    b = 0.2
    c = -65
    d = 2
    net = makeNeurons(total, startid, net, a, b, c, d)
    return net
}

func makeTCNeurons(total int, startid int, net[] Neuron) ([] Neuron) {
    var a, b, c, d float64
    a = 0.02
    b = 0.25
    c = -65
    d = 0.05
    net = makeNeurons(total, startid, net, a, b, c, d)
    return net
}

func makeRZNeurons(total int, startid int, net[] Neuron) ([] Neuron) {
    var a, b, c, d float64
    a = 0.1
    b = 0.25
    c = -65
    d = 2
    net = makeNeurons(total, startid, net, a, b, c, d)
    return net
}

func makeLTSNeurons(total int, startid int, net[] Neuron) ([] Neuron) {
    var a, b, c, d float64
    a = 0.02
    b = 0.25
    c = -65
    d = 2
    net = makeNeurons(total, startid, net, a, b, c, d)
    return net
}

func main() {
    var I[] float64
    var Istart int
    var Iend int
    var Ival float64

    var maxt int
    //var fired bool

    /*
        RS	./iz 500 0.02 0.2 -65 8
	IB	./iz 500 0.02 0.2 -55 4
	CH	./iz 500 0.02 0.2 -50 2
	FS	./iz 500 0.1 0.2 -65 2
	TC	./iz 500 0.02 0.25 -65 0.05
	RZ	./iz 500 0.1 0.25 -65 2
	LTS	./iz 500 0.02 0.25 -65 2
    */

    /*
    if len(os.Args) != 6 {
        fmt.Println(os.Args[0] + " maxt=500 a=0.02 b=0.2 c=-65 d=2")
        // max time cycles
        maxt = 500
        // recovery variable
        a = 0.02
        // sensitivy to a
        b = 0.2
        // after spike reset value
        c = -65
        // after spike reset of "a"
        d = 2
    } else {
        // max time cycles
        maxt, _ = strconv.Atoi(os.Args[1])
        // recovery variable
        a, _ = strconv.ParseFloat(os.Args[2], 64)
        // sensitivy to a
        b, _ = strconv.ParseFloat(os.Args[3], 64)
        // after spike reset value
        c, _ = strconv.ParseFloat(os.Args[4], 64)
        // after spike reset of "a"
        d, _ = strconv.ParseFloat(os.Args[5], 64)
    }
    */

    // set maxt to 500 ms
    maxt = 500
    // create current
    Istart = 10
    Iend = maxt
    Ival = 15.0
    I = makeCurrent(maxt, Istart, Iend, Ival)
    var net[] Neuron
    var synapses[] Synapse
    /*
    net = append(net, Neuron{a, b, c, d, u, v, false, I[0], 0, 0, 0})
    // 0.02 0.2 -50 2
    net = append(net, Neuron{a, b, -50, d, u, v, false, I[0], 0, 0, 1})
    net = append(net, Neuron{a, 0.25, -50, d, u, v, false, I[0], 0, 0, 2})
    */
    net = makeIBNeurons(4, 0, net)
    synapses = append(synapses, Synapse{0.5, 0, 1})
    synapses = append(synapses, Synapse{0.75, 0, 2})
    synapses = append(synapses, Synapse{0.25, 0, 3})
    synapses = append(synapses, Synapse{0.15, 1, 2})
    // [TODO]
    // make function to create network with x % of one type of neuron and y %
    // of another type
    // create synapse list which connects "layers" of neurons to the next
    // layer (ensuring feedforward connectionist architecture
    for t := 0; t < maxt; t++ {
        //fmt.Println(t, len(net))
        for i := range(net) {
            var n = net[i]
            //n.u, n.v, n.fired = neuron_iz(n.a, n.b, n.c, n.d, I[t], n.u, n.v)
            n.u, n.v, n.v_tick, n.fired = neuron_iz(n, I[t])
            net[i] = n
            fmt.Printf("%f,%f,%f,%v,%d,%d\n", n.u, n.v, I[t], n.fired, n.nodeid, t)
            if n.fired {
                fmt.Println("fired", i)
                // for now, reset I_tick
                n.I_tick = 0
                net[i] = n
                // update connected nodes
                for j := range(synapses) {
                    var s = synapses[j]
                    // get the synapse connected to the current neuron
                    if s.startnode == n.nodeid {
                        for k := range(net) {
                            var cn = net[k]
                            // find the neuron this synapse is connected to
                            if s.endnode == cn.nodeid {
                                // update the connected neurons I_tick value
                                cn.I_tick = cn.I_tick + (n.v_tick * s.weight)
                                fmt.Println("updated", cn.nodeid, "I_tick", cn.I_tick)
                                net[k] = cn
                            }
                        }
                    }
                }
            }
        }
        // once here, all firing have occured, we can clear I_tick
        // we assume no loops and a feed forward connectionist architecture
        for i := range(net) {
            var n = net[i]
            n.I_tick = 0
            net[i] = n
        }
        //u, v, fired = neuron_iz(a, b, c, d, I[t], u, v)
        //u = u_tick
        //v = v_tick
        //fmt.Printf("%f,%f,%f,%v\n", u, v, I[t], fired)
        //fmt.Printf("%f,%f,%f,%v,%d\n", n.u, n.v, I[t], n.fired, t)
    }
}
