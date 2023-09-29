import './App.css';
import React from 'react';
import Plot from 'react-plotly.js';

function App() {
  return (
    <div className="App">
      <Plot
        data={[
          {
            x: [1, 2, 3],
            y: [2, 6, 3],
            type: 'scatter',
            mode: 'lines+markers',
            marker: { color: 'red' },
          },
          { type: 'bar', x: [1, 2, 3], y: [2, 5, 3] },
        ]}
        layout={{ width: 1600, height: 800, title: 'Test plot' }}
      />
    </div>
  );
}

export default App;
