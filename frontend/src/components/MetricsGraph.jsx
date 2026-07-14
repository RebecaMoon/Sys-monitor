import {
  ResponsiveContainer,
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
} from "recharts";



function MetricsGraph({ title, history, dataKey }) {
  return (
    <div className="metric-graph">
      <h3>{title}</h3>

      <div className="graph-placeholder">
        <ResponsiveContainer width="100%" height="100%">
          <LineChart data={history}>
            <CartesianGrid
              stroke="rgba(72,242,251,0.08)"
              vertical={false}
            />

            <XAxis hide />

            <YAxis
              domain={[0, 100]}
              ticks={[0, 25, 50, 75, 100]}
              interval={0}
              tick={{ fill: '#4EA3BA', fontSize: 11 }}
              tickFormatter={(value) => `${value}%`}
              axisLine={false}
              tickLine={false}
            />

            <Line
              type="monotone"
              dataKey={dataKey}
              stroke="#48F2FB"
              strokeWidth={2}
              dot={false}
            />
          </LineChart>
        </ResponsiveContainer>
      </div>
    </div>
  )
}

export default MetricsGraph